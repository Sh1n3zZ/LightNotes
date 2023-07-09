package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	Token string `form:"token" binding:"required"`
}

func LoginAPI(c *gin.Context) {
	var form LoginForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "bad request",
		})
		return
	}
	access := Validate(form.Token)
	if access == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "user not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": access.Status,
	})
}
