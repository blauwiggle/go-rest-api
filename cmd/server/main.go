package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/blauwiggle/go-rest-api/internal/transport/http"
)

// App - the struct which contains things like pointers to the database, etc.
type App struct{}

// Run - sets up our application and starts the server
func (app *App) Run() error {
	fmt.Println("Starting up our REST API")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to start our REST API")
		return err
	}

	return nil
}

// App - the struct which contains things like pointers to the database, etc.
type App struct{}

// Run - sets up our application and starts the server
func (app *App) Run() error {
	fmt.Println("Starting up our REST API")
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
