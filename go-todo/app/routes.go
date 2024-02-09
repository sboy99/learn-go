package app

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
