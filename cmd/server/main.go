package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/PatriciaChebet/rest_api_with_go/internal/transport/http"
)

type App struct {
}

func (app *App) Run() error {
	fmt.Println("Setting our App UP")
	handler := transportHTTP.NewHandler()
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
