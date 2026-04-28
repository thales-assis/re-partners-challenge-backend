package httprouter

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	ProvideRouter,
)
