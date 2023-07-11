package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"strconv"
	"time"
)

const PaginationSize = 5

type LoginForm struct {
	Token string `form:"token" binding:"required"`
}

type NoteForm struct {
	Title string `form:"title" binding:"required"`
	Body  string `form:"body" binding:"required"`
}

type Note struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	UpdatedAt *time.Time `json:"updated_at"`
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
		code := GenerateCode(6)
		key := fmt.Sprintf("note:anymous:%s", code)
		res := cache.Get(c, key)
		if res.Err() != nil && len(res.Val()) == 0 {
			data, err := json.Marshal(form)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"status": false,
					"error":  "internal error",
				})
				return
			}
			cache.Set(c, key, data, time.Hour*24)
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
	res := cache.Get(c, fmt.Sprintf("note:anymous:%s", code))
	if res.Err() != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "code not found",
		})
		return
	}

	if res.Err() != nil || len(res.Val()) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "code not found",
		})
		return
	}

	var form NoteForm
	if err := json.Unmarshal([]byte(res.Val()), &form); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "internal error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"title":  form.Title,
		"body":   form.Body,
	})
}

func UserStateAPI(c *gin.Context) {
	username := c.MustGet("user").(string)
	c.JSON(http.StatusOK, gin.H{
		"status": len(username) != 0,
		"user":   username,
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
			"status":  false,
			"error":   "internal error",
			"message": err.Error(),
		})
		return
	}

	var noteID int
	if err := db.QueryRow("SELECT id FROM notes WHERE user_id = ? AND title = ? AND content = ? ORDER BY created_at DESC LIMIT 1", id, form.Title, form.Body).Scan(&noteID); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "internal error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"id":     noteID,
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

	var note Note
	var stamp []uint8
	if err := db.QueryRow("SELECT id, title, content, updated_at FROM notes WHERE user_id = ? AND id = ?", id, noteID).Scan(&note.ID, &note.Title, &note.Body, &stamp); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "note not found",
		})
		return
	}
	note.UpdatedAt = ConvertTime(stamp)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"note":   note,
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

	var id int
	if err := db.QueryRow("SELECT id FROM auth WHERE username = ?", username).Scan(&id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"error":  "user not found",
		})
		return
	}

	var total int
	if err := db.QueryRow("SELECT COUNT(*) FROM notes WHERE user_id = ?", id).Scan(&total); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"error":   "internal error",
			"message": err.Error(),
		})
		return
	}

	if total%PaginationSize > 0 {
		total = total/PaginationSize + 1
	} else {
		total = total / PaginationSize
	}

	if page > total || page < 1 {
		c.JSON(http.StatusOK, gin.H{
			"status":    true,
			"total":     total,
			"page":      page,
			"next_page": page+1 <= total,
			"prev_page": page-1 > 0,
			"notes":     make([]Note, 0),
		})
		return
	}

	rows, err := db.Query("SELECT id, title, content, updated_at FROM notes WHERE user_id = ? ORDER BY id DESC LIMIT ?, ?", id, (page-1)*PaginationSize, PaginationSize)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"error":   "internal error",
			"message": err.Error(),
		})
		return
	}
	defer rows.Close()

	notes := make([]Note, 0)
	for rows.Next() {
		var note Note
		var stamp []uint8
		if err := rows.Scan(&note.ID, &note.Title, &note.Body, &stamp); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"error":   "internal error",
				"message": err.Error(),
			})
			return
		}
		note.UpdatedAt = ConvertTime(stamp)
		notes = append(notes, note)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    true,
		"total":     total,
		"page":      page,
		"next_page": page+1 <= total,
		"prev_page": page-1 > 0,
		"notes":     notes,
	})
}
