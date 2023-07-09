package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func IsUserExist(db *sql.DB, username string) bool {
	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM auth WHERE username = ?", username).Scan(&count); err != nil {
		return false
	}
	return count > 0
}

func Login(c *gin.Context, token string) bool {
	user := Validate(token)
	if user == nil {
		return false
	}

	db := c.MustGet("db").(*sql.DB)
	if !IsUserExist(db, user.Username) {
		password := GenerateChar(64)
		_ = db.QueryRow("INSERT INTO auth (bind_id, username, token, password) VALUES (?, ?, ?, ?)",
			user.ID, user.Username, token, password)
		
	}
	return true
}
