package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sboy99/learn-go/go-todo/adapters/handlers/rest"
	"github.com/sboy99/learn-go/go-todo/adapters/repositories/postgres"
	"github.com/sboy99/learn-go/go-todo/infra/config"
	"github.com/sboy99/learn-go/go-todo/infra/database"
	"github.com/sboy99/learn-go/go-todo/internal/domain/models"
	"github.com/sboy99/learn-go/go-todo/internal/services"
	"gorm.io/gorm"
)

// ------------------------------STRUCT---------------------------------- //

type App struct {
	Router *fiber.App
	DB     *gorm.DB
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

	// create app
	app := &App{
		Router: fiber.New(),
		DB:     db,
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
	healtHandler := &rest.HealthHandler{
		Service: &services.HealthService{},
	}

	a.Router.Get("/", healtHandler.Check)
}

func (a *App) registerAuthRoutes() {
	userRepo := &postgres.UserPostgresRepository{
		DB: a.DB,
	}
	authService := &services.AuthService{
		UserRepo: userRepo,
	}

	authHandler := &rest.AuthHandler{
		Service: authService,
	}

	a.Router.Post("/auth/register", authHandler.Register)
}

func (a *App) registerTaskRoutes() {
	taskRepo := &postgres.TaskPostgresRepository{
		DB: a.DB,
	}
	taskService := &services.TaskService{
		TaskRepo: taskRepo,
	}
	taskHandler := &rest.TaskHandler{
		Service: taskService,
	}

	a.Router.Post("/tasks", taskHandler.Create)
}

func (a *App) registerUserRoutes() {
}
