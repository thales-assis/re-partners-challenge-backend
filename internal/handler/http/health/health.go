package health

import (
	"net/http"

	"github.com/re-partners-challenge-backend/internal/handler/http/viewmodel"
	"github.com/uptrace/bunrouter"
)

type HealthCheckHandler bunrouter.HandlerFunc

func ProvideHealthCheckHandler() HealthCheckHandler {
	return HealthHandler()
}

// HealthHandler godoc
// @Summary Check API health
// @Description Returns the current API health status.
// @Tags health
// @Produce json
// @Success 200 {object} viewmodel.HealthCheckResponse
// @Router /health [get]
func HealthHandler() HealthCheckHandler {
	return func(w http.ResponseWriter, r bunrouter.Request) error {
		return bunrouter.JSON(w, viewmodel.HealthCheckResponse{Status: "OK"})
	}
}
