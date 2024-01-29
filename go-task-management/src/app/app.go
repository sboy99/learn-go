package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sboy99/learn-go/go-task-management/src/adapters/controllers"
	"github.com/sboy99/learn-go/go-task-management/src/adapters/controllers/transformers"
	"github.com/sboy99/learn-go/go-task-management/src/adapters/repositories/postgres"
	"github.com/sboy99/learn-go/go-task-management/src/internal/domain/models"
	"github.com/sboy99/learn-go/go-task-management/src/internal/services"
	"github.com/sboy99/learn-go/go-task-management/src/routers"
)

type App struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func (app *App) Init() {
	// repositories
	userPostgresRepo := &postgres.UserPostgresRepository{
		DB: app.DB,
	}
	taskPostgresRepo := &postgres.TaskPostgresRepository{
		DB: app.DB,
	}

	// transformers
	userTransformer := &transformers.UserTransformer{}

	// services
	userService := &services.UserService{
		UserRepo: userPostgresRepo,
	}
	taskService := &services.TaskService{
		TaskRepo: taskPostgresRepo,
	}

	// controllers
	userController := &controllers.UserController{
		Service:     userService,
		Transformer: userTransformer,
	}
	taskController := &controllers.TaskController{
		Service: taskService,
	}

	// routers
	userRouter := &routers.UserRouter{
		Router:     app.Router,
		Controller: userController,
	}
	taskRouter := &routers.TaskRouter{
		Router:     app.Router,
		Controller: taskController,
	}

	// register routes
	userRouter.Register()
	taskRouter.Register()
}

func (app *App) RegisterRelations() {
	app.DB.AutoMigrate(&models.User{}, &models.Task{})
}

func (app *App) Start() {
	app.Router.Run("localhost:3000")
}
