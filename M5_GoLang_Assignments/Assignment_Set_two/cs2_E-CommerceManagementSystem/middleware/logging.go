package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware logs incoming requests and their responses
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		log.Printf("Started %s | %s\n", c.Request.Method, c.Request.URL.Path)

		c.Next()

		// Log request details
		log.Printf(
			"Completed %s | %s | %s | %s | %d | %s",
			c.ClientIP(),
			time.Now().Format(time.RFC3339),
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			time.Since(startTime),
		)
	}
}
