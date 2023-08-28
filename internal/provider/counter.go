package provider

import "sync"

var (
	queueIndexCounter int
	mutex             *sync.Mutex
)

func InitializeQueueIndexCounter() {
	queueIndexCounter = 0
	mutex = &sync.Mutex{}
}

func GetQueueName() string {
	mutex.Lock()
	defer mutex.Unlock()
	queueName := QueueNames[queueIndexCounter]
	queueIndexCounter = (queueIndexCounter + 1) % len(QueueNames)
	return queueName
}
