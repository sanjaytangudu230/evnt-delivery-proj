package internal

import (
	"context"
	"encoding/json"
	"eventDelivery/internal/provider"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushEvent(t *testing.T) {
	ctx := context.Background()

	// Assuming you have a running Redis server accessible via RedisClient
	// If you don't have RedisClient initialized, replace this with your setup code
	// var RedisClient *redis.Client
	queueName := "test-queue"

	// Set up test data
	testData := map[string]interface{}{
		"key": "value",
	}

	eventData, _ := json.Marshal(testData)

	// Call the function being tested
	provider.InitializeRedisClient()
	err := pushToQueue(ctx, eventData, queueName)

	// Assert that no error was returned from the function
	assert.NoError(t, err)

	// You might want to add additional assertions here to check if the data was pushed to the queue
}
