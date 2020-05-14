package tests

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/pressly/goose"
	"github.com/vsokoltsov/users-service/app"
)

const migrationsPath = "/app/app/migrations/"

// AppInstance saves App struct into variable for testing purposes
var AppInstance app.App

// MakeRequest performs test request on the application's routes
func MakeRequest(verb string, path string, params *bytes.Buffer) *httptest.ResponseRecorder {
	body := params
	if params == nil {
		body = bytes.NewBuffer([]byte{})
	}
	fmt.Println("PARAMS ARE ", body)
	req, _ := http.NewRequest(verb, path, body)
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
