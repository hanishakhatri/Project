package main

import "fmt"

// App - the struct which contain things like pointers
// to database connections
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting Up Our APP")
	return nil
}

func main() {
	fmt.Println("Go Rest API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
