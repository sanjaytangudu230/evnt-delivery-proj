package internal

import (
	"context"
	"encoding/json"
	"eventDelivery/internal/provider"
	"fmt"
)

func PushEvent(ctx context.Context, requestData map[string]interface{}) error {
	queueName := provider.GetQueueName()
	event, err := json.Marshal(requestData)
	if err != nil {
		fmt.Printf("Error occurred while converting to byte %s", err.Error())
		return err
	}
	return pushToQueue(ctx, event, queueName)
}

func pushToQueue(ctx context.Context, event []byte, queueName string) error {
	_, err := provider.RedisClient.RPush(ctx, queueName, event).Result()
	if err != nil {
		fmt.Println("Error pushing request to queue:", err)
		return err
	}
	fmt.Printf("Pushed request to %s: %s\n", queueName, event)
	return nil
}
