package httppackage

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	ProvideGetPackagesHandler,
	ProvideUpdatePackagesHandler,
	ProvideRouter,
)
