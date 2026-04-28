package health

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	ProvideHealthCheckHandler,
	ProvideRouter,
)
