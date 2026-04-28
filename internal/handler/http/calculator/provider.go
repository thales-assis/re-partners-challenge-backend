package calculator

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	ProvidePostCalculatorPackHandler,
	ProvideRouter,
)
