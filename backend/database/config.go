package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var dbsql *sql.DB

func DBInit() {
	var err error
	dbsql, err = sql.Open("postgres", os.Getenv("POSTGRE_URI"))
	if err != nil {
		panic(err)
	}
}

func GetDBConn() *sql.DB {
	return dbsql
}

func CloseDBConn() {
	if err := dbsql.Close(); err != nil {
		panic(err)
	}
}
