package main

import (
	"fmt"
	"net/http"

	"github.com/enriquesalceda/GoRestApi/internal/comment"
	"github.com/enriquesalceda/GoRestApi/internal/database"
	transportHTTP "github.com/enriquesalceda/GoRestApi/internal/transport/http"
	log "github.com/sirupsen/logrus"
)

// App:
// Contains application information
type App struct {
	Name    string
	Version string
}

func (app *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		},
	).Info("Setting up our APP")

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
		log.Error("Failed to setup the server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API")
	app := App{
		Name:    "Commenting service",
		Version: "1.0.0",
	}
	if err := app.Run(); err != nil {
		log.Error("Error starting up our REST API")
		log.Fatal(err)
	}
}
