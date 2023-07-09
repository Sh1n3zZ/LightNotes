// 该程序用于便签编写完成后保存时的处理

// 未来不在 我本该拥有海洋
// 孤岛极光 虚拟城邦 星球流浪
// 我将走向 无垠远方 我不遗忘 我的模样

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

var db *sql.DB
var cache *redis.Client

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	BindID   int64  `json:"bind_id"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type Response struct {
	Message string `json:"message"`
}

func main() {
	db = ConnectMySQL()
	cache = ConnectRedis()

	r := gin.Default()

	r.Use(AuthMiddleware()) // 添加中间件，用于验证用户登录状态

	r.POST("/save_note", saveNoteHandler)

	r.Run(fmt.Sprintf(":%d", viper.GetInt("server.port")))
}

func saveNoteHandler(c *gin.Context) {
	// 检查用户是否已登录
	user, ok := c.Get("user")
	if !ok {
		// 未登录，跳转回登录界面
		c.JSON(http.StatusUnauthorized, Response{Message: "未登录"})
		return
	}

	// 获取当前登录的用户ID和用户名
	username := user.(*User).Username

	// 获取用户输入的便签内容
	noteContent := c.PostForm("note_content")

	// 过滤用户输入的便签内容
	filteredNoteContent := sanitizeInput(noteContent)

	// 执行插入便签数据的 SQL 语句
	sql := "INSERT INTO notes (user_id, username, content) VALUES (?, ?, ?)"
	_, err := db.Exec(sql, user.(*User).ID, username, filteredNoteContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Message: "便签保存失败"})
		return
	}

	c.JSON(http.StatusOK, Response{Message: "便签保存成功"})
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
