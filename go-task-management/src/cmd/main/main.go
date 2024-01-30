package main

import (
	"github.com/sboy99/learn-go/go-task-management/src/app"
)

func main() {
	// create app
	taskManagementApp := app.Create()

	// register routes
	taskManagementApp.RegisterRouters()

	// initialize database relations
	taskManagementApp.RegisterRelations()

	// start the app
	taskManagementApp.Start()
}
