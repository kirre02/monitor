package main

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/kirre02/monitor-backend/internal/site/handler"
	"github.com/kirre02/monitor-backend/internal/site/service"
	"github.com/kirre02/monitor-backend/util"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
)

func main() {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config")
	}

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

	go runMigrations(config.DatabaseUrl, config.MigrationPath)

	log.Infof("Starting server at: %s", address)
	log.Fatal(srv.ListenAndServe())
}

func runMigrations(databaseURL, migrationPath string) error {
	db, err := initDB()
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	// Create a migration source
	m, err := migrate.NewWithDatabaseInstance(migrationPath, "postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func initDB() (*sqlx.DB, error) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config")
	}

	uri := config.DatabaseUrl
	db, err := sqlx.Connect("postgres", uri)
	if err != nil {
		return nil, fmt.Errorf("not able to connect to postgres: \nuri: %s\nerror:%v", uri, err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("unable to ping postgres \nuri: %s\nerror:%v", uri, err)
	}

	return db, nil
}
