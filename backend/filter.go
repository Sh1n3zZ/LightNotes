// 该程序用于过滤 HTML 标签以防止 XSS 注入

// 君の笑い方はなぜか淋しさに似てた
// 君の歌い方は今日の朝焼けに見えた
// 何千年後の人類が何をしているかより
// まだ誰も知らない顔で　笑う君を見たい

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	FilteredInput string `json:"filtered_input"`
}

func main() {
	r := gin.Default()

	r.Use(AuthMiddleware()) // 添加中间件，用于验证用户登录状态

	r.POST("/filter", filterHandler)

	r.Run(fmt.Sprintf(":%d", viper.GetInt("server.port")))
}

func filterHandler(c *gin.Context) {
	// 检查用户是否已登录
	_, ok := c.Get("user")
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	input := c.PostForm("input")
	filteredInput := sanitizeInput(input)

	response := Response{
		FilteredInput: filteredInput,
	}

	c.JSON(http.StatusOK, response)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user := ParseToken(c, token)
		if user == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
