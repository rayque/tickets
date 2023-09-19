package server

import (
	"github.com/gin-gonic/gin"
	"tickets/internal/handlers/observability"
)

type RoutesController struct {
	observabilityController *observability.Controller
}

func NewRoutesController(
	observabilityController *observability.Controller,
) *RoutesController {
	return &RoutesController{observabilityController: observabilityController}
}

func (r RoutesController) Define(server *gin.Engine) {
	server.GET("/check", r.observabilityController.HealthCheck)
}
