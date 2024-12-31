package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InputValidationMiddleware(requiredFields []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyBytes, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read request body"})
			c.Abort()
			return
		}

		var payload map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
			c.Abort()
			return
		}

		for _, field := range requiredFields {
			if _, ok := payload[field]; !ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required field: " + field})
				c.Abort()
				return
			}
		}

		// Reset the body for the next handler
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		c.Next()
	}
}
