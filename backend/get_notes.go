// 该程序用于获取登录用户所属账号下的便签文件及其页数

// 後悔ばっかが募って
// でも後戻りはもうできなくて
// ずっと　ずっと
// "君だらけ"の毎日

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strconv"
)

type Note struct {
	Content string `json:"content"`
}

type Response struct {
	Notes      []Note `json:"notes"`
	TotalPages int    `json:"totalPages"`
}

var db *sql.DB

func main() {
	db = ConnectMySQL()

	r := gin.Default()

	r.GET("/get_notes", getNotesHandler)

	r.Run(fmt.Sprintf(":%d", viper.GetInt("server.port")))
}

func getNotesHandler(c *gin.Context) {
	// 检查用户是否已登录
	_, ok := c.Get("user_id")
	if !ok {
		// 未登录，返回空数组
		c.JSON(http.StatusOK, Response{})
		return
	}

	// 获取当前登录的用户ID
	userID := c.GetInt64("user_id")

	// 获取页码和每页便签数量
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("perPage", "5"))
	offset := (page - 1) * perPage

	// 查询当前用户的便签数量
	countSql := fmt.Sprintf("SELECT COUNT(*) as total FROM notes WHERE user_id = %d", userID)
	var totalCount int
	err := db.QueryRow(countSql).Scan(&totalCount)
	if err != nil {
		log.Println("Failed to get note count:", err)
		c.JSON(http.StatusInternalServerError, Response{})
		return
	}

	// 计算总页数
	totalPages := int((totalCount + perPage - 1) / perPage)

	// 查询当前页的便签数据
	sql := fmt.Sprintf("SELECT content FROM notes WHERE user_id = %d ORDER BY id DESC LIMIT %d, %d", userID, offset, perPage)
	rows, err := db.Query(sql)
	if err != nil {
		log.Println("Failed to get notes:", err)
		c.JSON(http.StatusInternalServerError, Response{})
		return
	}
	defer rows.Close()

	notes := make([]Note, 0)
	for rows.Next() {
		var note Note
		err := rows.Scan(&note.Content)
		if err != nil {
			log.Println("Failed to scan note:", err)
			c.JSON(http.StatusInternalServerError, Response{})
			return
		}
		notes = append(notes, note)
	}

	response := Response{
		Notes:      notes,
		TotalPages: totalPages,
	}

	c.JSON(http.StatusOK, response)
}

func ConnectMySQL() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.db"),
	))
	if err != nil {
		log.Fatalln("Failed to connect to MySQL server:", err)
	}
	log.Println("Connected to MySQL server successfully")

	CreateNotesTable(db)

	return db
}

func CreateNotesTable(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS notes (
		  id INT PRIMARY KEY AUTO_INCREMENT,
		  user_id INT NOT NULL,
		  content VARCHAR(255) NOT NULL,
		  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		  FOREIGN KEY (user_id) REFERENCES auth(id)
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}
