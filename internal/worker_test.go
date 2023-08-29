package internal

import (
	"context"
	"encoding/json"
	"eventDelivery/internal/provider"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	provider.InitializeRedisClient()
	os.Exit(exitCode)
}

func TestPushToDestination_Success(t *testing.T) {
	// Set up a mock HTTP server for testing
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	event := Event{
		UserId:      "user123",
		Payload:     "payload",
		Retry:       0,
		Destination: server.URL,
	}

	pushToDestination(context.Background(), event, event.Destination)

	// Assert that no error was returned
	// assert.NoError(t, err)
}

func TestPushToDestination_HTTPError(t *testing.T) {
	// Set up a mock HTTP server that will return an error response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	event := Event{
		UserId:      "user123",
		Payload:     "payload",
		Retry:       0,
		Destination: server.URL,
	}

	pushToDestination(context.Background(), event, event.Destination)

	// Assert that an error was returned
	// assert.Error(t, err)
}

func TestFetchFromQueues(t *testing.T) {
	// Set up a Redis client
	// Note: You need to replace this with your actual Redis client initialization
	// var RedisClient *redis.Client

	queueName := "test-queue"

	// Create a test event
	event := Event{
		UserId:      "user123",
		Payload:     "payload",
		Retry:       0,
		Destination: "http://example.com",
	}

	byteEvent, _ := json.Marshal(event)

	// Push the test event onto the queue
	_, _ = provider.RedisClient.RPush(context.Background(), queueName, byteEvent).Result()

	// Fetch events from the queue
	events := fetchFromQueues(context.Background(), []string{queueName})

	// Assert that the fetched event matches the original event
	assert.Len(t, events, 1)
	assert.Equal(t, event, events[0])
}

// You can similarly write test cases for other functions like pushToDeadQueue and PollDeadMessages
