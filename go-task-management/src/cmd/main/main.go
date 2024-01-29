package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sboy99/learn-go/go-task-management/src/app"
	"github.com/sboy99/learn-go/go-task-management/src/infra/db"
)

func main() {
	// get router
	router := gin.Default()

	// connect to database
	db := db.Connect()

	// create app
	app := &app.App{
		Router: router,
		DB:     db,
	}

	// initialize routes
	app.Init()

	// initialize database relations
	app.RegisterRelations()

	// start the app
	app.Start()
}
