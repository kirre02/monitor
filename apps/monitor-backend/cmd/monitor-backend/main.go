package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

func initDB() (*sqlx.DB, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get environment variables
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	sslMode := os.Getenv("POSTGRES_SSLMODE")

	// Build connection string
	uri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbName, sslMode)

	db, err := sqlx.Connect("postgres", uri)
	if err != nil {
		return nil, fmt.Errorf("not able to connect to postgres: \nuri: %s\nerror:%v", uri, err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("unable to ping postgres \nuri: %s\nerror:%v", uri, err)
	}

	return db, nil
}
