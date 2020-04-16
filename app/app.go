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
	Router           *mux.Router
	ConnectionString string
}

// Initialize populates App struct with
// necessary parameters
func (app *App) Initialize(env string) {
	dbConName := utils.GetDatabaseConnection(env)
	app.ConnectionString = os.Getenv(dbConName)
	utils.InitDB(app.ConnectionString)
	app.Router = api.InitRouter(mux.NewRouter())
	app.Router.Use(loggingMiddleware)
}

// Start run application
func (app *App) Start() {
	log.Println("Starting server on port 8000")
	err := http.ListenAndServe(":8000", app.Router)
	if err != nil {
		log.Fatalln(err)
	}
}

// loggingMiddleware defines necessary logging level for app
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
