package main

import (
	"fmt"
	"net/http"

	"github.com/enriquesalceda/GoRestApi/internal/comment"
	"github.com/enriquesalceda/GoRestApi/internal/database"
	transportHTTP "github.com/enriquesalceda/GoRestApi/internal/transport/http"
)

// App:
// The struct which contains things like
// pointers to DB connections
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up our APP")

	var err error
	db, err := database.NewDatabase()

	if err != nil {
		return err
	}

	err = database.MigrateDB(db)

	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
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
