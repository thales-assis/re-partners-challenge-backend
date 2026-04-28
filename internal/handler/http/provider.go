package http

import (
	"github.com/google/wire"
	"github.com/re-partners-challenge-backend/internal/handler/http/calculator"
	"github.com/re-partners-challenge-backend/internal/handler/http/health"
	"github.com/re-partners-challenge-backend/internal/handler/http/pack"
)

var ProviderSet = wire.NewSet(
	health.ProviderSet,
	calculator.ProviderSet,
	pack.ProviderSet,
)
