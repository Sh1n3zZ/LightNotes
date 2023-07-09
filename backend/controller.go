package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"strconv"
	"time"
)

const PaginationSize = 10

type LoginForm struct {
	Token string `form:"token" binding:"required"`
}

type NoteForm struct {
	Title string `form:"title" binding:"required"`
	Body  string `form:"body" binding:"required"`
}

type Note struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
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
	var form NoteForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "bad request",
		})
		return
	}

	if len(form.Title) > 120 || len(form.Body) > 10240 {
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
	var form NoteForm
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

func UserSaveAPI(c *gin.Context) {
	var form NoteForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "bad request",
		})
		return
	}

	if len(form.Title) > 120 || len(form.Body) > 4000 {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "title or body too long",
		})
		return
	}

	username := c.MustGet("user").(string)

	if len(username) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "user not found",
		})
		return
	}

	db := c.MustGet("db").(*sql.DB)
	var id int
	if err := db.QueryRow("SELECT id FROM auth WHERE username = ?", username).Scan(&id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "user not found",
		})
		return
	}
	_, err := db.Exec("INSERT INTO notes (user_id, title, content) VALUES (?, ?, ?)", id, form.Title, form.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "internal error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

func UserGetAPI(c *gin.Context) {
	username := c.MustGet("user").(string)
	noteID := c.Query("id")

	if len(username) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "user not found",
		})
		return
	}

	db := c.MustGet("db").(*sql.DB)
	var id int
	if err := db.QueryRow("SELECT id FROM auth WHERE username = ?", username).Scan(&id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "user not found",
		})
		return
	}

	var title, body string
	if err := db.QueryRow("SELECT title, content FROM notes WHERE user_id = ? AND id = ?", id, noteID).Scan(&title, &body); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "note not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"title":  title,
		"body":   body,
	})
}

func UserUpdateAPI(c *gin.Context) {
	var form NoteForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "bad request",
		})
		return
	}

	if len(form.Title) > 120 || len(form.Body) > 4000 {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "title or body too long",
		})
		return
	}

	username := c.MustGet("user").(string)
	noteID := c.Query("id")

	if len(username) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "user not found",
		})
		return
	}

	db := c.MustGet("db").(*sql.DB)
	var id int
	if err := db.QueryRow("SELECT id FROM auth WHERE username = ?", username).Scan(&id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "user not found",
		})
		return
	}

	_, err := db.Exec("UPDATE notes SET title = ?, content = ? WHERE user_id = ? AND id = ?", form.Title, form.Body, id, noteID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "internal error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

func UserDeleteAPI(c *gin.Context) {
	username := c.MustGet("user").(string)
	noteID := c.Query("id")

	if len(username) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "user not found",
		})
		return
	}

	db := c.MustGet("db").(*sql.DB)
	var id int
	if err := db.QueryRow("SELECT id FROM auth WHERE username = ?", username).Scan(&id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "user not found",
		})
		return
	}

	_, err := db.Exec("DELETE FROM notes WHERE user_id = ? AND id = ?", id, noteID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "internal error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

func UserListAPI(c *gin.Context) {
	// pagination
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "bad request",
		})
	}

	username := c.MustGet("user").(string)

	if len(username) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "user not found",
		})
		return
	}

	db := c.MustGet("db").(*sql.DB)
	var total int
	if err := db.QueryRow("SELECT COUNT(*) FROM notes").Scan(&total); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "internal error",
		})
		return
	}

	if total%PaginationSize > 0 {
		total = total/PaginationSize + 1
	} else {
		total = total / PaginationSize
	}

	if page > total {
		page = total
	}

	var id int
	if err := db.QueryRow("SELECT id FROM auth WHERE username = ?", username).Scan(&id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "user not found",
		})
		return
	}

	rows, err := db.Query("SELECT id, title, content FROM notes WHERE user_id = ? ORDER BY id DESC LIMIT ?, ?", id, (page-1)*PaginationSize, PaginationSize)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "internal error",
		})
		return
	}
	defer rows.Close()

	notes := make([]Note, 0)
	for rows.Next() {
		var note Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Body); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"error":  "internal error",
			})
			return
		}
		notes = append(notes, note)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    true,
		"total":     total,
		"page":      page,
		"next_page": page+1 <= total,
		"prev_page": page-1 >= 0,
		"notes":     notes,
	})
}
