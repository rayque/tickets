package server

import "github.com/gin-gonic/gin"

type GinWebServers struct {
	Controllers *RoutesController
}

func NewGinServer(routesController *RoutesController) *GinWebServers {
	return &GinWebServers{routesController}
}

type GinWebServerInit func()

func (g *GinWebServers) CreateServer() (*gin.Engine, GinWebServerInit) {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	g.Controllers.Define(server)
	return server, func() {
		server.Run(":8080")
	}
}
