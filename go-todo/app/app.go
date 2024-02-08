package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sboy99/learn-go/go-todo/infra/config"
	"github.com/sboy99/learn-go/go-todo/infra/database"
	"github.com/sboy99/learn-go/go-todo/internal/domain/models"
	"gorm.io/gorm"
)

// ------------------------------STRUCT---------------------------------- //

type App struct {
	Router   *fiber.App
	DB       *gorm.DB
	Handlers *RestHandlers
}

// ------------------------------PUBLIC---------------------------------- //

func Create() (*App, error) {
	// load environment variables
	err := config.LoadEnv()
	if err != nil {
		return nil, err
	}

	// connect to database
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	// create handler
	handler := &Handler{
		DB: db,
	}

	// create app
	app := &App{
		Router:   fiber.New(),
		DB:       db,
		Handlers: handler.initRestHandlers(),
	}

	return app, nil
}

func (a *App) RegisterMiddlewares() {
	a.Router.Use(logger.New())
	a.Router.Use(recover.New())
}

func (a *App) RegisterRoutes() {
	a.registerHealthRoutes()
	a.registerAuthRoutes()
	a.registerTaskRoutes()
}

func (a *App) RegisterDatabaseRelations() {
	a.DB.AutoMigrate(&models.User{}, &models.Task{})
}

func (a *App) Serve() {
	port := config.HTTP_PORT
	err := a.Router.Listen(port)
	log.Fatal(err.Error())
}

// ------------------------------PRIVATE---------------------------------- //

func (a *App) registerHealthRoutes() {
	a.Router.Get("/", a.Handlers.HealthHandler.Check)
}

func (a *App) registerAuthRoutes() {
	a.Router.Post("/auth/register", a.Handlers.AuthHandler.Register)
	a.Router.Post("/auth/login", a.Handlers.AuthHandler.Login)
}

func (a *App) registerTaskRoutes() {
	a.Router.Post("/tasks", a.Handlers.TaskHandler.Create)
}

func (a *App) registerUserRoutes() {
}
