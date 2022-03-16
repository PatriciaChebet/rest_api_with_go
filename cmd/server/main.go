package main

import (
	"fmt"
	"net/http"

	"github.com/PatriciaChebet/rest_api_with_go/internal/comment"
	"github.com/PatriciaChebet/rest_api_with_go/internal/database"
	transportHTTP "github.com/PatriciaChebet/rest_api_with_go/internal/transport/http"
)

type App struct {
}

func (app *App) Run() error {
	fmt.Println("Setting our App UP")

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
	handler.SetUpRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Go REST API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting App")
		fmt.Println(err)
	}
}
