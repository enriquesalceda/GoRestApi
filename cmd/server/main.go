package main

import "fmt"

// App:
// The struct which contains things like
// pointers to DB connections
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up our APP")
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
