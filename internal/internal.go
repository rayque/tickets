package internal

import (
	"go.uber.org/fx"
	"tickets/internal/handlers"
	"tickets/internal/server"
)

var Module = fx.Options(
	server.Module,
	handlers.Module,
)
