package internal

import (
	"github.com/google/wire"

	"github.com/re-partners-challenge-backend/internal/handler"
	"github.com/re-partners-challenge-backend/internal/infra"
)

var ProviderSet = wire.NewSet(
	infra.ProviderSet,
	handler.ProviderSet,
)
