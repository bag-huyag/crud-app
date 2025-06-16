package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {

	// connStr := "host=postgres user=postgres password=postgres dbname=users_db sslmode=disable"

	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	// host := "localhost"
	// port := "5432"
	user := "postgres"
	password := "postgres"
	dbname := "users_db"
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to PostgreSQL!")
}
