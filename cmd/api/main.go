package main

import (
	"go.uber.org/fx"
	"tickets/internal"
	"tickets/internal/server"
)

func main() {

	fx.New(
		internal.Module,
		fx.Invoke(bootstrap),
	)
}

func bootstrap(server *server.GinWebServers) {
	_, runServer := server.CreateServer()
	runServer()
}
