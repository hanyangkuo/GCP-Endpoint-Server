package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"google.golang.org/api/option"
	"log"
)

func main() {
	publish("golang-pratice", "my-topic", "Hello World")
}


func publish(projectID, topicID, msg string) error {
	// projectID := "my-project-id"
	// topicID := "my-topic"
	// msg := "Hello World"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID, option.WithCredentialsFile("D:\\Users\\tw014\\go\\src\\github.com\\Osquery-Endpoint-Server\\golang-pratice-44b21bdbc0f0.json"))
	if err != nil {
		log.Fatalf("pubsub.NewClient: %v", err)
	}

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		log.Fatalf("Get: %v", err)
	}
	log.Printf("Published a message; msg ID: %v\n", id)
	return nil
}