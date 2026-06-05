package database

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
)

var DB *sql.DB 

func Connect() {
	var err error
	DB, err = sql.Open("sqlite", "iot.db")
	if err != nil {
		log.Fatal(err)
	}
	
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database successfully!")
}