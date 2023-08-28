package server

import (
	"context"
	"eventDelivery/api"
	"eventDelivery/internal"
	"eventDelivery/internal/provider"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

func Start(ctx context.Context) {

	router := gin.Default()

	registerRoutes(router)
	configRuntime()
	provider.InitializeRedisClient()
	provider.InitializeQueueIndexCounter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go internal.PollMessages(ctx)
	go internal.PollDeadMessages(ctx)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("Server start failed: %v\n", err)
	}

	err = server.Close()

	if err != nil {
		fmt.Println("Server shutdown failed")
	}
}

func configRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
}

func registerRoutes(router *gin.Engine) {
	group := router.Group("/api/event-delivery")
	{
		api.HealthCheck(group)
		api.IngestEvent(group)
	}
}
