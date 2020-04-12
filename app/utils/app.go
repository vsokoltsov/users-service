package utils

import (
	"log"
	"net/http"
)

const developmentDBConString = "DB_CON"
const testDBString = "DB_CON_TEST"

// LoggingMiddleware defines necessary logging level for app
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
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
