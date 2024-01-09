package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/kirre02/monitor-backend/internal/site/handler"
	"github.com/kirre02/monitor-backend/internal/site/service"
	_ "github.com/lib/pq"
)

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	siteSvc := service.NewSiteService(db)
	siteHandler := &handler.SiteHandler{
		Service: siteSvc,
	}

	router := chi.NewRouter()

	router.Mount("/api/v1", siteHandler.Routes())

	address := "0.0.0.0:9090"

	srv := &http.Server{
		Addr:              address,
		Handler:           router,
		ReadTimeout:       5 * time.Second, // Adjust the timeouts as needed
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       15 * time.Second,
	}

	log.Infof("Starting server at: %s", address)
	log.Fatal(srv.ListenAndServe())
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
