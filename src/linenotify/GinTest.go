package linenotify

import (
	"github.com/gin-gonic/gin"
)

func GinApi() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func GetValue() string {
	return "test"
}
