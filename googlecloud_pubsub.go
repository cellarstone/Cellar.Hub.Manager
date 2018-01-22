package main

import (
	"fmt"

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

	fmt.Println("TEST11")

	// Creates a client.
	client, err := pubsub.NewClient(ctx, googleCloudProjectID)
	if err != nil {
		fmt.Printf("Failed to create client: %v", err)
	}

	fmt.Println("TEST22")

	// Sets the name for the new topic.
	googleCloudPubsubTopic = "cellarDevice_" + cellarDeviceID

	topic := client.Topic(googleCloudPubsubTopic)
	ok, err := topic.Exists(ctx)
	if err != nil {
		// TODO: Handle error.
		fmt.Printf("Failed to check if topic exists: %v", err)
	}
	if !ok {

		fmt.Println("TEST33")

		// Creates the new topic.
		topic, err := client.CreateTopic(ctx, googleCloudPubsubTopic)
		if err != nil {
			fmt.Printf("Failed to create topic: %v", err)
		}

		fmt.Println("TEST44")

		fmt.Printf("Topic %v created.\n", topic)
	}

}

func Publish(message string) {
	ctx := context.Background()

	fmt.Println("TEST55")

	client, err := pubsub.NewClient(ctx, googleCloudProjectID)
	if err != nil {
		fmt.Printf("Failed to create client: %v", err)
	}

	fmt.Println("TEST66")

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
			fmt.Printf("Failed publish to Cloud PubSub: %v", err)
		}
		fmt.Printf("Published a message with a message ID: %s\n", id)
	}

	fmt.Println("TEST77")
}
