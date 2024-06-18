package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var err error

func DBConnection() {
	DB, err = sql.Open("mysql", "kanban:bw1qJGj@tcp(127.0.0.1:6000)/kanban?multiStatements=true")

	if err != nil {
		panic(err.Error())
	}

	if err = DB.Ping(); err != nil {
		panic("Failed to ping database: " + err.Error())
	} else {
		log.Println("Successfully pinged database")
	}
}
