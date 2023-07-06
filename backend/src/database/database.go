package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectToDB() (db *sql.DB, err error) {
	db, err = sql.Open("postgres", config.DBConn)
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return
	}

	return
}
