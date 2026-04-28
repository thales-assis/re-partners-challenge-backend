package app

import (
	"github.com/google/wire"
	"github.com/re-partners-challenge-backend/internal"
)

var ProviderSet = wire.NewSet(
	internal.ProviderSet,
	ProvideApplication,
)
