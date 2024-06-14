package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB
var err error

func DBConnection() {
	DB, err = sql.Open("mysql", "kanban:bw1qJGj@tcp(127.0.0.1:6000)/kanban?multiStatements=true")

	if err != nil {
		panic(err.Error())
	}

	// defer db.Close()

	if err = DB.Ping(); err != nil {
		panic("Failed to ping database: " + err.Error())
	} else {
		log.Println("Succesfully pinged database")
	}
}
