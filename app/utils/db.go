package utils

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

// InitDB initializes db instance
func InitDB(dataSource string) {
	var err error
	DB, err = sqlx.Connect("postgres", dataSource)
	if err != nil {
		log.Panic(err)
	}

	if err = DB.Ping(); err != nil {
		log.Panic(err)
	}
}
