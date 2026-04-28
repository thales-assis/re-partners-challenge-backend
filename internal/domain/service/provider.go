package service

import (
	"github.com/google/wire"

	calculatorservice "github.com/re-partners-challenge-backend/internal/domain/service/calculator"
	packservice "github.com/re-partners-challenge-backend/internal/domain/service/pack"
)

var ProviderSet = wire.NewSet(
	calculatorservice.ProviderSet,
	packservice.ProviderSet,
)
