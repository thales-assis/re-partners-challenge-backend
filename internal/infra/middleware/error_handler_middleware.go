package middleware

import (
	"net/http"

	"github.com/re-partners-challenge-backend/internal/infra/log"
	"github.com/uptrace/bunrouter"
)

type Middleware bunrouter.MiddlewareFunc

func Handler(handlers ...ErrorHandlerFunc) Middleware {
	return func(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
		return func(w http.ResponseWriter, req bunrouter.Request) error {

			err := next(w, req)
			if err == nil {
				return nil
			}

			for _, handler := range handlers {
				httpErrorResponse, ok := handler(err)(w, req)
				if ok {
					w.WriteHeader(httpErrorResponse.Status)
					return bunrouter.JSON(w, httpErrorResponse)
				}
			}

			return err
		}
	}
}

func ProvideErrorHandlerMiddleware(logger *log.ZapLogger) Middleware {
	return Handler(
		UnprocessableEntityErrorHandler(),
		InternalServerErrorHandler(logger),
	)
}
