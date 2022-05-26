package main

import (
	"fmt"
	"net/http"

	"github.com/blauwiggle/go-rest-api/internal/comment"
	"github.com/blauwiggle/go-rest-api/internal/database"
	transportHTTP "github.com/blauwiggle/go-rest-api/internal/transport/http"
)

// App - the struct which contains things like pointers to the database, etc.
type App struct{}

// Run - sets up our application and starts the server
func (app *App) Run() error {
	fmt.Println("Starting up our REST API")

	var err error
	db, err := database.NewDatabase()

	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		fmt.Println("Failed to migrate database")
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to start our REST API")
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
