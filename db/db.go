package db

import (
	"database/sql"
	"log"
)

const SqliteDBDsn = "./flickr.db"

func NewSQLiteConnection() (*sql.DB, error) {
	dbConn, err := sql.Open("sqlite3", SqliteDBDsn)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return dbConn, nil
}
