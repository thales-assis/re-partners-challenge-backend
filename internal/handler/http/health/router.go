package health

import (
	"github.com/uptrace/bunrouter"
)

type Router struct {
	healthCheckHandler HealthCheckHandler
}

func ProvideRouter(healthCheckHandler HealthCheckHandler) Router {
	return Router{
		healthCheckHandler: healthCheckHandler,
	}
}

func (r Router) Register(router *bunrouter.Group) {
	router.GET("/health", bunrouter.HandlerFunc(r.healthCheckHandler))
}
