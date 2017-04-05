package shim

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

type httpRequest struct {
	Body       string            `json:"body"`
	Header     map[string]string `json:"headers"`
	Method     string            `json:"method"`
	RemoteAddr string            `json:"remote_addr"`
	URL        string            `json:"url"`
}

type httpResponse struct {
	Body       string            `json:"body"`
	Header     map[string]string `json:"headers"`
	StatusCode int               `json:"status_code"`
	Error      string            `json:"error,omitempty"`
}

//ServeHTTP takes your http.HandlerFunc and runs it through the shim.
func ServeHTTP(h http.HandlerFunc) {
	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		SendHTTPError(500, err)
		return
	}

	var request httpRequest
	err = json.Unmarshal(stdin, &request)
	if err != nil {
		SendHTTPError(500, err)
		return
	}

	r, err := http.NewRequest(request.Method, request.URL, bytes.NewBufferString(request.Body))
	if err != nil {
		SendHTTPError(500, err)
		return
	}

	for k, v := range request.Header {
		r.Header.Add(k, v)
	}

	r.RemoteAddr = request.RemoteAddr

	if r.URL.Path == "" {
		r.URL.Path = "/"
	}

	w := httptest.NewRecorder()

	h(w, r)

	resp := w.Result()
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		SendHTTPError(500, err)
		return
	}

	// base64 encode the body
	base64body := base64.StdEncoding.EncodeToString(body)

	header := make(map[string]string)
	for k, v := range resp.Header {
		header[k] = strings.Join(v, ",")
	}
	response := httpResponse{
		Body:       base64body,
		Header:     header,
		StatusCode: resp.StatusCode,
	}

	out, err := json.Marshal(&response)
	if err != nil {
		SendHTTPError(500, err)
		return
	}

	_, err = bytes.NewReader(out).WriteTo(os.Stdout)
	if err != nil {
		SendHTTPError(500, err)
		return
	}
}

func SendHTTPError(code int, err error) {
	response := httpResponse{
		Error:      err.Error(),
		StatusCode: code,
	}

	out, err := json.Marshal(&response)
	if err != nil {
		out = []byte("{\"error\": \"" + err.Error() + "\"}")
	}

	os.Stdout.Write(out)
}
