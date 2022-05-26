package main

import "fmt"

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
