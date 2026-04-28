package calculatorservice

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	ProvideCalculatorService,
)
