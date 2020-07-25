package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func DbConn() (db *sql.DB){
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "starwars"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Fatal(500, err)
        return
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	return db
}
