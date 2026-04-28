package packusecase

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	ProvidePackUseCase,
)
