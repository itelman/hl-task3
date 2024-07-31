package config

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
)

func OpenDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//psqlInfo1 := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", "localhost", "5432", "admin", "password", "database")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	// Read the SQL file
	filePath := "./migrations/postgres/00001_initial.up.sql"
	sqlBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read SQL file: %v\n", err)
	}
	sqlString := string(sqlBytes)

	// Execute the SQL commands
	_, err = db.Exec(sqlString)
	if err != nil {
		log.Fatalf("Failed to execute SQL commands: %v\n", err)
	}

	return db, nil
}
