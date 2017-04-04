package shim

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
)

type PubsubHandlerFunc func(message pubsub.Message)

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

	h(message)
}
