package http

import (
	"github.com/google/wire"
	"github.com/re-partners-challenge-backend/internal/handler/http/health"
)

var ProviderSet = wire.NewSet(
	health.ProviderSet,
)
