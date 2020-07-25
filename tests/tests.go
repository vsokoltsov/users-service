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
	app "github.com/vsokoltsov/users-service/pkg"
	"github.com/vsokoltsov/users-service/pkg/utils"
)

const migrationsPath = "/app/pkg/migrations/"

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

// // clearDB deletes all the created records from database
func clearDB() {
	tx := utils.DB.MustBegin()
	tx.MustExec("delete from users")
	tx.Commit()
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
	clearDB()
	defer clearDB()
	defer db.Close()

	errDB := goose.Run("up", db, migrationsPath)
	if errDB != nil {
		log.Fatal(errDB)
	}
	exitVal := m.Run()
	os.Exit(exitVal)
}
