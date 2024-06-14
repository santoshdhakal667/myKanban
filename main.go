package main

import (
	"fmt"
	"log"
	"os"

	"example.com/kanban/database"
	"example.com/kanban/handlers"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

// const migrationsDir = "database/migration"

// //go:embed database/migration/*.sql
// var MigrationsFS embed.FS

// func checkErr(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func connectDatabase() error {
// 	db, err := sql.Open("sqlite3", "./database/db.sqlite")
// 	if err != nil {
// 		return err
// 	}

// 	model.DB = db
// 	return nil
// }

// var DB *sql.DB

// func initDatabase() {
// 	connectionStr := "kanban:bw1qJGj@tcp(127.0.0.1:6000)/kanban?multiStatements=true"
// 	var err error
// 	DB, err = sql.Open("mysql", connectionStr)
// 	if err != nil {
// 		log.Fatalf("Failed to open database connection: %v", err)
// 	}
// 	err = DB.Ping()
// 	if err != nil {
// 		log.Fatalf("Failed to ping database: %v", err)
// 	}
// }

func main() {
	database.DBConnection()
	// initDatabase()
	// // --- (1) ----
	// // Recover Migrator
	// migrator := migrator.MustGetNewMigrator(MigrationsFS, migrationsDir)

	// // --- (2.1) ----
	// // Get the DB instance
	// connectionStr := "kanban:bw1qJGj@tcp(127.0.0.1:6000)/kanban?multiStatements=true"
	// conn, err := sql.Open("mysql", connectionStr)
	// if err != nil {
	// 	panic(err)
	// }
	// defer conn.Close()

	// // --- (2) ----
	// // Apply migrations
	// err = migrator.ApplyMigrations(conn)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Migrations applied!!\n")

	LOG_FILE := "./log/output.log"
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/test", func(ctx *gin.Context) {
			fmt.Println("lol")
		})
		v1.GET("/boards", func(c *gin.Context) {
			handlers.GetBoard(c)
		})
		// v1.GET("/boards/:id", func(c *gin.Context) {
		// 	handlers.GetBoardByID(c)
		// })
		// v1.POST("/boards", func(c *gin.Context) {
		// 	handlers.PostBoard(c)
		// })
		// v1.PUT("/boards/:id", func(c *gin.Context) {
		// 	handlers.PutBoard(c)
		// })
		// v1.DELETE("/boards/:id", func(c *gin.Context) {
		// 	handlers.DeleteBoardByID(c)
		// })
	}

	router.Run("localhost:9000")
}
