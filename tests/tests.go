package tests

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/pressly/goose"
	"github.com/vsokoltsov/users-service/app"
)

const migrationsPath = "/app/app/migrations/"

var AppInstance app.App

func MakeRequest(verb string, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(verb, path, nil)
	rr := httptest.NewRecorder()
	AppInstance.Router.ServeHTTP(rr, req)
	return rr
}

// TestMain shows defaults TestMain for test
func TestMain(m *testing.M) {
	appEnv := os.Getenv("APP_ENV")
	AppInstance = app.App{}
	AppInstance.Initialize(appEnv)

	db, err := sql.Open("postgres",
		AppInstance.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	errDB := goose.Run("up", db, migrationsPath)
	if errDB != nil {
		log.Fatal(errDB)
	}
	exitVal := m.Run()
	os.Exit(exitVal)
}
