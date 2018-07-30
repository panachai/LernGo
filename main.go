package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// GetValue()

	writeConsoleLog()

	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", GinApi)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	router.Run(":8080")
}

func writeConsoleLog() {
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
}
