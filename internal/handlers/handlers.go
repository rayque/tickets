package handlers

import (
	"go.uber.org/fx"
	"tickets/internal/handlers/observability"
)

var Module = fx.Options(
	fx.Provide(observability.NewObservabilityController))
