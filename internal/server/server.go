package server

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewGinServer),
	fx.Provide(NewRoutesController),
)
