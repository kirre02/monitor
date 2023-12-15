package main

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"
    _"github.com/lib/pq"
)

func main() {
    db, err := initDB("")
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()
}


func initDB(uri string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", uri)
	if err != nil {
        return nil, fmt.Errorf("not able to connect to postgres: \nuri: %s\nerror:%v", uri, err)
	}

    if err = db.Ping(); err != nil {
        return nil, fmt.Errorf("Unable to ping postgres \nuri: %s\nerror:%v", uri, err)
    }

	return db, nil
}
