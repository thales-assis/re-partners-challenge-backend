package http

import (
	"github.com/google/wire"
	"github.com/re-partners-challenge-backend/internal/handler/http/health"
	"github.com/re-partners-challenge-backend/internal/handler/http/httppackage"
)

var ProviderSet = wire.NewSet(
	health.ProviderSet,
	httppackage.ProviderSet,
)
