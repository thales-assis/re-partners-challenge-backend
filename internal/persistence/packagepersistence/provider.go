package packagepersistence

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	ProvidePackageRepository,
)
