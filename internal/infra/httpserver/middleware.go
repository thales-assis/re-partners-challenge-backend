package httpserver

import (
	"github.com/re-partners-challenge-backend/internal/infra/log"
	"github.com/re-partners-challenge-backend/internal/infra/middleware"
	"github.com/uptrace/bunrouter"
)

func ProvideCoreMiddlewares(
	logger *log.ZapLogger,
	errorHandlerMiddleware middleware.Middleware,
) []bunrouter.MiddlewareFunc {
	return []bunrouter.MiddlewareFunc{
		bunrouter.MiddlewareFunc(errorHandlerMiddleware),
	}
}
