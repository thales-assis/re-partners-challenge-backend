package persistence

import (
	"github.com/google/wire"
	"github.com/re-partners-challenge-backend/internal/persistence/database"
	"github.com/re-partners-challenge-backend/internal/persistence/packagepersistence"
)

var ProviderSet = wire.NewSet(
	database.ProviderSet,
	packagepersistence.ProviderSet,
)
