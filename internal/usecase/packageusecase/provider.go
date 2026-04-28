package packageusecase

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	ProvidePackageUseCase,
)
