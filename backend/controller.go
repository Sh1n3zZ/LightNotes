package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"time"
)

type LoginForm struct {
	Token string `form:"token" binding:"required"`
}

type AnonymousNoteForm struct {
	Title string `form:"title" binding:"required"`
	Body  string `form:"body" binding:"required"`
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

	state, token := Login(c, form.Token)
	if !state {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "user not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"token":  token,
	})
}

func AnonymousSendAPI(c *gin.Context) {
	var form AnonymousNoteForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "bad request",
		})
		return
	}

	if len(form.Title) > 120 || len(form.Body) > 1024*10 {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "title or body too long",
		})
		return
	}

	cache := c.MustGet("cache").(*redis.Client)
	for {
		code := fmt.Sprintf("note:anymous:%s", GenerateCode(8))
		res := cache.Get(c, code)
		if res.Err() != nil && len(res.Val()) == 0 {
			cache.Set(c, code, form, time.Hour*24)
			c.JSON(http.StatusOK, gin.H{
				"status": true,
				"code":   code,
			})
			return
		}
	}
}

func AnonymousGetAPI(c *gin.Context) {
	code := c.Query("code")
	if len(code) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "bad request",
		})
		return
	}

	cache := c.MustGet("cache").(*redis.Client)
	res := cache.Get(c, code)
	if res.Err() != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "code not found",
		})
		return
	}
	var form AnonymousNoteForm
	if err := res.Scan(&form); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "code not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"title":  form.Title,
		"body":   form.Body,
	})
}
