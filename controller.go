package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginAPI(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(400, gin.H{
			"error": "token is required",
		})
		return
	}
	access := Validate(token)
	c.JSON(http.StatusOK, gin.H{
		"status": access,
	})
}
