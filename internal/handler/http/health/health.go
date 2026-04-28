package health

import (
	"net/http"

	"github.com/uptrace/bunrouter"
)

type HealthCheckHandler bunrouter.HandlerFunc

func ProvideHealthCheckHandler() HealthCheckHandler {
	return HealthHandler()
}

func HealthHandler() HealthCheckHandler {
	return func(w http.ResponseWriter, r bunrouter.Request) error {
		return bunrouter.JSON(w, bunrouter.H{
			"status": "OK",
		})
	}
}
