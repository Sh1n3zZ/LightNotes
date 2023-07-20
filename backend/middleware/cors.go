package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var allowedOrigins = []string{
	"https://fystart.cn",
	"https://www.fystart.cn",
	"https://40code.com",
	"https://www.40code.com",
}

func Contains[T comparable](value T, slice []T) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if Contains(origin, allowedOrigins) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

			if c.Request.Method == "OPTIONS" {
				c.Writer.Header().Set("Access-Control-Max-Age", "3600")
				c.AbortWithStatus(http.StatusOK)
				return
			}
		}

		c.Next()
	}
}
