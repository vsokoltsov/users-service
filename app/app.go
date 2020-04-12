package app

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/vsokoltsov/users-service/app/api"
	"github.com/vsokoltsov/users-service/app/utils"
)

// App represents application
// with configs and pointers to necessary structures
type App struct {
	Router *mux.Router
}

// Initialize populates App struct with
// necessary parameters
func (app *App) Initialize(env string) {
	dbConName := utils.GetDatabaseConnection(env)
	dbCon := os.Getenv(dbConName)
	utils.InitDB(dbCon)
	app.Router = api.InitRouter(mux.NewRouter())
	app.Router.Use(utils.LoggingMiddleware)
	log.Println("Starting server on port 8000")
}

// Start run application
func (app *App) Start() {
	err := http.ListenAndServe(":8000", app.Router)
	if err != nil {
		log.Fatalln(err)
	}
}
