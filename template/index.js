require('dotenv').config()

function shim (data, callback) {
  let responsePayload = ''

  // Spawn the function and inject the env from the parent process.
  const p = require('child_process').execFile('./{{.FunctionName}}', [], {
    env: process.env,
    // Large responses over stdin were overruning the default stdio buffer.
    // Since we don't need to work about any other concurrent requests or tenants, give it a healthy max
    maxBuffer: 1048576 * 25
  })

  p.stdin.setEncoding('utf-8')

  // Log standard err messages to standard out
  p.stderr.on('data', data => {
    console.log(data)
  })

  p.stdout.on('data', data => {
    responsePayload += data.toString()
  })

  p.on('error', err => {
    console.log('go-cloud-fn shim received error', err)
  })

  p.on('close', exitCode => {
    if (exitCode > 0) {
      return callback(new Error('Shim exited unsuccessfully with code=' + exitCode), responsePayload)
    }
    return callback(null, responsePayload)
  })

  // Write the input data to stdin
  p.stdin.write(JSON.stringify(data))
  p.stdin.end()
}

// Handle http request
function handleHttp (req, res) {
  var requestBody

  switch (req.get('content-type')) {
    case 'application/json':
      requestBody = JSON.stringify(req.body)
      break
    case 'application/x-www-form-urlencoded':
      // The body parser for cloud functions does this, so just play along
      // with it, sorry man! Maybe we should construct some kind of proper
      // form request body? or not. let's keep it this way for now, as
      // This is how cloud functions behaves.
      req.set('content-type', 'application/json')
      requestBody = JSON.stringify(req.body)
      break
    case 'application/octet-stream':
      requestBody = req.body
      break
    case 'text/plain':
      requestBody = req.body
      break
  }

  var fullUrl = req.protocol + '://' + req.get('host') + req.originalUrl

  var httpRequest = {
    'body': requestBody,
    'headers': req.headers,
    'method': req.method,
    'remote_addr': req.ip,
    'url': fullUrl
  }

  shim(httpRequest, (err, result) => {
    if (err) {
      err.code = 500
      throw err
    }

    let data = null
    try {
      data = JSON.parse(result)
    } catch (ex) {
      err = new Error('Failed to parse shim response: ' + ex)
      err.code = 500
      throw err
    }

    res.status(data.status_code)
    res.set(data.headers)
    res.send(Buffer.from(data.body, 'base64'))
  })
}

// {{ if .TriggerHTTP }}
exports['{{.FunctionName}}'] = function (req, res) {
  return handleHttp(req, res)
}
// {{ else }}
exports['{{.FunctionName}}'] = function (event, callback) {
  shim(event.data, (err, result) => {
    if (err) {
      return callback(err)
    }
    if (result && result.error) {
      return callback(new Error(result.error))
    }
    callback()
  })
}
// {{ end }}