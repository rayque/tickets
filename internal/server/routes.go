package server

import (
	"github.com/gin-gonic/gin"
	"tickets/internal/handlers"
)

func DefineRoutes(engine *gin.Engine, handlers *handlers.HttpHandlers) {
	engine.GET("/check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	engine.GET("/users/:uuid", handlers.UserHandler.GetUser)
	engine.POST("/users", handlers.UserHandler.CreateUser)
	engine.PUT("/users/:uuid", handlers.UserHandler.UpdateUser)
	engine.DELETE("/users/:uuid", handlers.UserHandler.DeleteUser)
}
