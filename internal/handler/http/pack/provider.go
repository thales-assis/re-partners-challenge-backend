package pack

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	ProvideGetPacksHandler,
	ProvideUpdatePacksHandler,
	ProvideRouter,
)
