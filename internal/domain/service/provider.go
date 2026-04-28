package service

import (
	"github.com/google/wire"

	packageservice "github.com/re-partners-challenge-backend/internal/domain/service/package"
)

var ProviderSet = wire.NewSet(
	packageservice.ProviderSet,
)
