package usecase

import (
	"github.com/google/wire"
	"github.com/re-partners-challenge-backend/internal/usecase/packageusecase"
)

var ProviderSet = wire.NewSet(
	packageusecase.ProviderSet,
)
