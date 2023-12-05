package main

import (
	"context"
	"fmt"
	"tema1/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Printf("Failed to start application: %s", err)
	}
}