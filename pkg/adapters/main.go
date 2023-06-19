package adapters

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

func GetDB() *sqlx.DB {
	if db == nil {
		newDb, err := sqlx.Open("sqlite3", os.Getenv("DATABASE_URL"))

		if err != nil {
			println(err.Error())
			os.Exit(1)
			return nil
		}

		err = newDb.Ping()
		if err != nil {
			println(err.Error())
			os.Exit(1)
			return nil
		}

		db = newDb
	}

	return db
}
