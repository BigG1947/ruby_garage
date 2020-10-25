package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
)

var db *sql.DB

func InitDB() {
	conn, err := sql.Open("sqlite3", "projects.db")
	if err != nil {
		log.Fatalf("%s\n", err)
		return
	}

	if err := conn.Ping(); err != nil {
		log.Fatalf("%s\n", err)
		return
	}

	db = conn
}

func InitTestDB() {
	conn, err := sql.Open("sqlite3", "projects_test.db")
	if err != nil {
		log.Fatalf("%s\n", err)
		return
	}

	if err := conn.Ping(); err != nil {
		log.Fatalf("%s\n", err)
		return
	}

	sqlTestData, err := ioutil.ReadFile("../test_data.sql")
	if err != nil {
		log.Fatalf("%s\n", err)
		return
	}

	if _, err := conn.Exec(string(sqlTestData)); err != nil {
		log.Fatalf("%s\n", err)
		return
	}

	db = conn
}
