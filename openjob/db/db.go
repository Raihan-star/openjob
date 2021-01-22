package db

import (
	"database/sql"
	"fmt"
	"openjob/config"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	connectionString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.DB_HOST, conf.DB_PORT, conf.DB_USERNAME, conf.DB_HOST, conf.DB_NAME)

	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic("connectionString error")
	}
	/*
		err = db.Ping()
		if err != nil {
			panic("DSN Invalid")
		}
	*/
	defer db.Close()

}

func CreateCon() *sql.DB {
	return db
}
