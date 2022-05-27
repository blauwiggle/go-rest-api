package main

import (
	"net/http"

	"github.com/blauwiggle/go-rest-api/internal/comment"
	"github.com/blauwiggle/go-rest-api/internal/database"
	transportHTTP "github.com/blauwiggle/go-rest-api/internal/transport/http"

	log "github.com/sirupsen/logrus"
)

// App - contain application information
type App struct {
	Name    string
	Version string
}

// Run - sets up our application and starts the server
func (app *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(log.Fields{
		"AppName":    app.Name,
		"AppVersion": app.Version,
	}).Info("Setting up our application")

	var err error
	db, err := database.NewDatabase()

	if err != nil {
		log.Error("Failed to connect to database")
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		log.Error("Failed to migrate database")
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Error("Failed to start our REST API")
		return err
	}

	return nil
}

func main() {
	app := App{
		Name:    "Go Commenting REST API",
		Version: "1.0.0",
	}

	if err := app.Run(); err != nil {
		log.Error("Error starting up our REST API")
		log.Fatal(err)
	}

}
