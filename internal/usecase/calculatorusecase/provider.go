package calculatorusecase

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	ProvideCalculatorUseCase,
)
