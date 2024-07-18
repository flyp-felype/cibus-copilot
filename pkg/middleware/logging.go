package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func LoggingMidleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Request: %s %s ", c.Request.Method, c.Request.URL)
	}
}
