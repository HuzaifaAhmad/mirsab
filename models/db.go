package models

import (
	"database/sql"
	"fmt"
	"os"

	//Psql Driver
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = os.Getenv("DBPASS")
	dbname   = "mirsab"
)

var DB *sql.DB

func init() {
	var err error

	stmt := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	DB, err = sql.Open("postgres", stmt)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}
}
