package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = ""
	port     = 7313
	user     = ""
	password = ""
	dbname   = ""
)

func InitDB() (db *sql.DB, err error) {
	connstr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	db, err = sql.Open("postgres", connstr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic("cannot connect to database.. ping")
	}
	fmt.Println("Successfully connect to database")
	return
}
