package usecase

import (
	"github.com/google/wire"
	"github.com/re-partners-challenge-backend/internal/usecase/calculatorusecase"
	"github.com/re-partners-challenge-backend/internal/usecase/packusecase"
)

var ProviderSet = wire.NewSet(
	calculatorusecase.ProviderSet,
	packusecase.ProviderSet,
)
