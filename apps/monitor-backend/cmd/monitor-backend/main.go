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
	"github.com/kirre02/monitor-backend/internal/check"
	"github.com/kirre02/monitor-backend/internal/site/handler"
	"github.com/kirre02/monitor-backend/internal/site/service"
	"github.com/kirre02/monitor-backend/util"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
)

func main() {
	config, err := util.LoadConfig(".")
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

	checkSvc := check.NewCheckService(db)
	checkHandler := &check.CheckHandler{
		Svc: checkSvc,
	}

	router := chi.NewRouter()

	router.Mount("/api/v1/site", siteHandler.Routes())
	router.Mount("/api/v1/check", checkHandler.Routes())

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

	// Whenever the main function exits, stop the checkservice
	checkSvc.StopCron()
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
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Errorf("cannot load config: %v", err)
		return nil, err
	}

	uri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresPassword, config.PostgresDb, config.PostgresSslmode)

	db, err := sqlx.Open("postgres", uri)
	if err != nil {
		return nil, fmt.Errorf("not able to connect to postgres: \nuri: %s\nerror:%v", uri, err)
	}

	if err = db.Ping(); err != nil {
		db.Close() // Close the database connection if Ping fails
		return nil, fmt.Errorf("unable to ping postgres \nuri: %s\nerror:%v", uri, err)
	}

	return db, nil
}
