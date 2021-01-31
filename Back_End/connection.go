package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

// not finish singleton connection
func getDB(db *sql.DB, err error)  {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "1234"
	dbName := "golangSimpleApplication"
	db, err = sql.Open(dbDriver,dbUser+":"+dbPass+"@/"+dbName)
	return
}
