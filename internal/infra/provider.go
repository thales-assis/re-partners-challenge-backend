package infra

import (
	"github.com/google/wire"

	"github.com/re-partners-challenge-backend/internal/infra/config"
	"github.com/re-partners-challenge-backend/internal/infra/httprouter"
	"github.com/re-partners-challenge-backend/internal/infra/httpserver"
	"github.com/re-partners-challenge-backend/internal/infra/log"
	"github.com/re-partners-challenge-backend/internal/infra/middleware"
)

var ProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	middleware.ProviderSet,
	httprouter.ProviderSet,
	httpserver.ProviderSet,
)
