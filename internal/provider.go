package internal

import (
	"github.com/google/wire"

	"github.com/re-partners-challenge-backend/internal/domain/service"
	"github.com/re-partners-challenge-backend/internal/handler"
	"github.com/re-partners-challenge-backend/internal/infra"
	"github.com/re-partners-challenge-backend/internal/persistence"
	"github.com/re-partners-challenge-backend/internal/usecase"
)

var ProviderSet = wire.NewSet(
	infra.ProviderSet,
	persistence.ProviderSet,
	service.ProviderSet,
	usecase.ProviderSet,
	handler.ProviderSet,
)
