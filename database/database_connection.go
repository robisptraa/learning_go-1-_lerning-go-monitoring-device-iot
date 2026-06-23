package database

import (
	"database/sql"
	"log"
	_"modernc.org/sqlite"
	"io/ioutil"
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

	log.Println("Connected to database successfully")

	schema, err := ioutil.ReadFile("database/iot.sql")
	if err != nil {
		log.Fatal("Failed to read iot.sql:", err)
	}

	_, err = DB.Exec(string(schema))
	if err != nil {
		log.Fatal("Failed to initialize tables:", err)
	}
}