package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	"example.com/kanban/handlers"
	model "example.com/kanban/model"

	"github.com/gin-gonic/gin"

	"github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func connectDatabase() error {
	db, err := sql.Open("sqlite3", "./database/db.sqlite")
	if err != nil {
		return err
	}

	model.DB = db
	return nil
}

var db *sql.DB

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "kanban",
		Passwd: "bw1qJGj",
		// User:   os.Getenv("DBUSER"),
		// Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "kanban",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	LOG_FILE := "./log/output.log"
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// err = connectDatabase()
	// checkErr(err)

	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/boards", func(c *gin.Context) {
			handlers.GetBoard(c)
		})
		v1.GET("/cards", func(c *gin.Context) {
			handlers.GetBoard(c)
		})
		v1.GET("/boards/:id", func(c *gin.Context) {
			handlers.GetBoardByID(c)
		})
		v1.POST("/boards", func(c *gin.Context) {
			handlers.PostBoard(c)
		})
		v1.PUT("/boards/:id", func(c *gin.Context) {
			handlers.PutBoard(c)
		})
		v1.DELETE("/boards/:id", func(c *gin.Context) {
			handlers.DeleteBoardByID(c)
		})

	}

	router.Run("localhost:9090")
}
