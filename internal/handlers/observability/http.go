package observability

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
}

func NewObservabilityController() *Controller {
	return &Controller{}
}

type observability struct {
	Status bool `json:"status"`
}

func (o *Controller) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, observability{
		Status: true,
	})
}
