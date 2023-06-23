package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

var (
	dsn        = flag.String("dsn", os.Getenv("DATABASE_URL"), "datasource name")
	migrations = flag.String("migrations", os.Getenv("MIGRATIONS"), "migrations directory")
	down       = flag.Int("down", 0, "number of migrations to rollback")
)

func main() {
	flag.Parse()
	if *dsn == "" {
		log.Fatal("dsn is required")
		return
	} else if *migrations == "" {
		log.Fatal("migrations is required")
		return
	}

	db, err := sql.Open("sqlite3", *dsn)
	if err != nil {
		log.Fatal(err)
		return
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", *migrations), "sqlite3", driver)
	if err != nil {
		log.Fatal(err)
		return
	}
	
	if *down > 0 {
		err = m.Steps(0 - *down)
	} else {
		err = m.Up()
	}

	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
		return
	}
}
