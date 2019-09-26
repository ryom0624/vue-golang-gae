package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var Conn *sql.DB

func init() {
	conn, err := sql.Open("sqlite3", "./db/articles.sqlite3")
	if err != nil {
		panic(err)
	}
	Conn = conn
}
