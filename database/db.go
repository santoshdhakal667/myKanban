package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

const (
	username = "root"
	password = "supersecret"
	hostname = "127.0.0.1:6000"
	dbname   = "kanban"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", dsn("kanban"))
	if err != nil {
		log.Fatalf("error %s when opening DB", err)
	}

	DB.SetConnMaxLifetime(5 * time.Minute)
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	if err = DB.PingContext(ctx); err != nil {
		panic("Failed to ping database: " + err.Error())
	} else {
		log.Println("Successfully pinged database")
	}
}