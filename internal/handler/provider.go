package handler

import (
	"github.com/re-partners-challenge-backend/internal/handler/http"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	http.ProviderSet,
)
