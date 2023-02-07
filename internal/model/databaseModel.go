package model

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func InitDb() error {
	connectDbStr := "user=postgres password=postgres host=localhost dbname=postgres port=5432  sslmode=disable"
	var err error
	Db, err = sql.Open("postgres", connectDbStr)
	if err != nil {
		log.Println(err)
	}
	//defer db.Close() // why this need ?
	err = Db.Ping()
	if err != nil {
		log.Println(err)
	}
	return err
}
