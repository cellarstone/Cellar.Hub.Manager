package main

import (
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
)

// CLOUD PUBSUB ->
// PUBLISH      - send info about device (cellarDeviceID, cellarHostName, cellarMACaddress,  Wifi, IP ... etc.)

func sendDeviceInfo() {
	CreateTopicIfNotExist()

	Publish(cellarDeviceInfo)
}

//************************************************************
//************************************************************
//************************************************************
//************************************************************
// GOOGLE CLOUD PUBSUB

func CreateTopicIfNotExist() {
	ctx := context.Background()

	// Creates a client.
	client, err := pubsub.NewClient(ctx, googleCloudProjectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the name for the new topic.
	googleCloudPubsubTopic = "cellarDevice_" + cellarDeviceID

	topic := client.Topic(googleCloudPubsubTopic)
	ok, err := topic.Exists(ctx)
	if err != nil {
		// TODO: Handle error.
		log.Fatalf("Failed to check if topic exists: %v", err)
	}
	if !ok {
		// Creates the new topic.
		topic, err := client.CreateTopic(ctx, googleCloudPubsubTopic)
		if err != nil {
			log.Fatalf("Failed to create topic: %v", err)
		}

		fmt.Printf("Topic %v created.\n", topic)
	}

}

func Publish(message string) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, googleCloudProjectID)
	if err != nil {
		// TODO: Handle error.
	}

	topic := client.Topic(googleCloudPubsubTopic)
	defer topic.Stop()
	var results []*pubsub.PublishResult
	r := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(message),
	})
	results = append(results, r)
	// Do other work ...
	for _, r := range results {
		id, err := r.Get(ctx)
		if err != nil {
			// TODO: Handle error.
		}
		fmt.Printf("Published a message with a message ID: %s\n", id)
	}
}
