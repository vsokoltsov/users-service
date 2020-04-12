package main

import (
	"os"

	"github.com/vsokoltsov/users-service/app"
)

func main() {
	appEnv := os.Getenv("APP_ENV")
	app := app.App{}
	app.Initialize(appEnv)
	app.Start()
}
