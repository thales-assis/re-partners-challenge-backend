package packservice

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	ProvidePackService,
)
