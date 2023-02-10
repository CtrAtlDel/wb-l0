package model

import (
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Db *sql.DB

func InitDb() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
	usr := os.Getenv("Username")
	psw := os.Getenv("Password")
	port := os.Getenv("Port")
	host := os.Getenv("Host")
	dbname := os.Getenv("Dbname")
	connectDbStr := "user=" + usr + " password=" + psw + " host=" + host + " dbname=" + dbname + " port=" + port + " sslmode=disable"
	Db, err = sql.Open("postgres", connectDbStr)
	if err != nil {
		log.Println(err)
	}
	err = Db.Ping()
	if err != nil {
		log.Println(err)
	}
	return err
}
