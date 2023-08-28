package api

import (
	"encoding/json"
	"eventDelivery/internal"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func HealthCheck(router *gin.RouterGroup) {
	router.GET("health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"status":  http.StatusOK,
			"message": "Request Successful",
			"data":    "Server UP",
		})
		return
	})
}

func IngestEvent(router *gin.RouterGroup) {
	router.POST("ingest-event", func(c *gin.Context) {
		requestBody, err := getRequestBody(c)
		err = internal.PushEvent(c, requestBody)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"status":  http.StatusOK,
				"message": "Request Successful",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
			})
		}
		return
	})
}

func getRequestBody(c *gin.Context) (map[string]interface{}, error) {
	resultBody := map[string]interface{}{}

	if c.Request.Body == nil {
		return resultBody, nil
	}
	jsonBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return resultBody, err
	}

	err = json.Unmarshal(jsonBody, &resultBody)
	return resultBody, err
}
