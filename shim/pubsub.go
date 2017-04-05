package shim

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
)

type PubsubHandlerFunc func(message pubsub.Message) error

type PubsubResult struct {
	Error error `json:"error"`
}

func HandlePubSubMessage(h PubsubHandlerFunc) {
	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var message pubsub.Message
	err = json.Unmarshal(stdin, &message)
	if err != nil {
		log.Fatal(err)
	}

	result := PubsubResult{
		Error: h(message),
	}
	out, err := json.Marshal(&result)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(out)
}
