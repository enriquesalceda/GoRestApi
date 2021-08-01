package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/enriquesalceda/GoRestApi/internal/transport/http"
)

// App:
// The struct which contains things like
// pointers to DB connections
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up our APP")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup the server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
