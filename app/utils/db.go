package utils

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	DB *sqlx.DB
)

const (
	developmentDBConString = "DB_CON"
	testDBString           = "DB_CON_TEST"
)

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

// GetDatabaseConnection returns name of the
// conection string based on env variable value
func GetDatabaseConnection(env string) string {
	switch env {
	case "development":
		return developmentDBConString
	case "test":
		return testDBString
	default:
		return developmentDBConString
	}
}
