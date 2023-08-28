package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"eventDelivery/internal/provider"
	"fmt"
	"net/http"
	"time"
)

type Event struct {
	UserId      string `json:"user_id"`
	Payload     string `json:"payload"`
	Retry       int    `json:"retry"`
	Destination string `json:"destination"`
}

func PollMessages(ctx context.Context) {
	for {
		events := fetchFromQueues(ctx, provider.QueueNames)
		for _, event := range events {
			pushToDestinations(ctx, event)
		}
	}
}

func PollDeadMessages(ctx context.Context) {
	for {
		var deadLetterQueue []string
		deadLetterQueue = append(deadLetterQueue, provider.DeadLetterQueue)
		events := fetchFromQueues(ctx, deadLetterQueue)
		for _, event := range events {
			pushToDestination(ctx, event, event.Destination)
		}
	}
}

func pushToDestinations(ctx context.Context, event Event) {
	for _, destination := range provider.Destinations {
		go pushToDestination(ctx, event, destination)
	}
}

func pushToDestination(ctx context.Context, event Event, destination string) {
	jsonBytes, err := json.Marshal(event)
	if err != nil {
		fmt.Println("Error while marshalling: ", err)
		return
	}
	req, err := http.NewRequest(http.MethodPost, destination, bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		event.Destination = destination
		event.Retry = event.Retry + 1
		err = pushToDeadQueue(ctx, event)
		if err != nil {
			fmt.Println("Error while pushing to dead letter queue: ", err)
		}
		return
	}

}

func pushToDeadQueue(ctx context.Context, event Event) error {
	if event.Retry == 4 {
		fmt.Println("Retry limit reached. Discarding the event: ", event)
		return nil
	}
	byteEvent, err := json.Marshal(event)
	if err != nil {
		fmt.Println("Error occurred while converting to byte ", err.Error())
		return err
	}
	_, err = provider.RedisClient.RPush(ctx, provider.DeadLetterQueue, byteEvent).Result()
	if err != nil {
		fmt.Println("Error pushing request to queue:", err)
		return err
	}
	return nil
}

func fetchFromQueues(ctx context.Context, queues []string) []Event {
	var events []Event
	for _, queueName := range queues {
		var event Event
		message, err := provider.RedisClient.BRPop(ctx, 1*time.Second, queueName).Result()
		if err != nil {
			continue
		}

		err = json.Unmarshal([]byte(message[1]), &event)
		if err != nil {
			fmt.Println("Error occurred while unmarshalling ", err.Error())
			continue
		}
		events = append(events, event)
	}
	return events
}
