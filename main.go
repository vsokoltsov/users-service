package main

import (
	"os"

	app "github.com/vsokoltsov/users-service/pkg"
)

func main() {
	appEnv := os.Getenv("APP_ENV")
	app := app.App{}
	app.Initialize(appEnv)
	app.Start()
}
