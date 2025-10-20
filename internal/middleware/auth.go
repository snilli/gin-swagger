package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Auth is a simple authentication middleware
// In production, this would validate JWT tokens or API keys
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Example: Check for API key in header
		apiKey := c.GetHeader("X-API-Key")

		// For demo purposes, accept any non-empty API key
		// In production, validate against a database or secret
		if apiKey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized - missing API key",
			})
			c.Abort()
			return
		}

		// Set user info in context for later use
		c.Set("api_key", apiKey)
		c.Next()
	}
}
