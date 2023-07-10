package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
)

func ConnectMySQL() *sql.DB {
	// connect to MySQL
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.db"),
	))
	if err != nil {
		log.Fatalln("Failed to connect to MySQL server: ", err)
	} else {
		log.Println("Connected to MySQL server successfully")
	}

	CreateUserTable(db)
	CreateNoteTable(db)
	return db
}

func CreateUserTable(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS auth (
		  id INT PRIMARY KEY AUTO_INCREMENT,
		  bind_id INT UNIQUE,
		  username VARCHAR(24) UNIQUE,
		  token VARCHAR(255) NOT NULL,
		  password VARCHAR(64) NOT NULL
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateNoteTable(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS notes (
		  id INT PRIMARY KEY AUTO_INCREMENT,
		  user_id INT NOT NULL,
		  title VARCHAR(120) NOT NULL,
		  content TEXT(40000) NOT NULL,
		  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		  FOREIGN KEY (user_id) REFERENCES notes.auth(id)
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectRedis() *redis.Client {
	// connect to redis
	rds := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
	_, err := rds.Ping(context.Background()).Result()

	if err != nil {
		log.Fatalln("Failed to connect to Redis server: ", err)
	} else {
		log.Println("Connected to Redis server successfully")
	}

	if viper.GetBool("debug") {
		rds.FlushAll(context.Background())
		log.Println("Flushed all cache")
	}

	return rds
}
