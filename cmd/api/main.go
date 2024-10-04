package main

import (
	"github.com/gin-gonic/gin"
	"tickets/internal/application/usecases"
	"tickets/internal/handlers"
	"tickets/internal/repositories"
	"tickets/internal/server"
)

func main() {
	engine := gin.Default()

	httpHandlers := buildHttpHandlers()
	server.DefineRoutes(engine, httpHandlers)

	err := engine.Run()
	if err != nil {
		return
	}
}

func buildHttpHandlers() *handlers.HttpHandlers {
	userRepository := repositories.NewUserRepository()
	useUseCase := usecases.NewUserUseCase(userRepository)

	UserHandler := handlers.NewUserHandler(useUseCase)
	return handlers.NewHttpHandlers(*UserHandler)
}
