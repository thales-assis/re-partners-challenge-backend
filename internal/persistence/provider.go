package persistence

import (
	"github.com/google/wire"
	"github.com/re-partners-challenge-backend/internal/persistence/database"
	"github.com/re-partners-challenge-backend/internal/persistence/packpersistence"
)

var ProviderSet = wire.NewSet(
	database.ProviderSet,
	packpersistence.ProviderSet,
)
