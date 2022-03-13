package main

import "fmt"

type App struct {
}

func (app *App) Run() error {
	fmt.Println("Setting our App UP")
	return nil
}

func main() {
	fmt.Println("Go REST API")
	app := App{}
    if err := app.Run(); if err != nil{
		fmt.Println("Error starting App")
		fmt.Println(err)
	}
}
