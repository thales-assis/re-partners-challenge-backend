package middleware

import (
	"net/http"

	"github.com/re-partners-challenge-backend/internal/handler/http/viewmodel"
	"github.com/re-partners-challenge-backend/internal/infra/log"
	"github.com/uptrace/bunrouter"
)

const (
	internalServerErrorMessage string = "the server encountered an error and could not complete your request."
)

func InternalServerErrorHandler(logger *log.ZapLogger) ErrorHandlerFunc {
	return func(err error) func(w http.ResponseWriter, req bunrouter.Request) (viewmodel.HttpErrorResponse, bool) {
		return func(w http.ResponseWriter, req bunrouter.Request) (viewmodel.HttpErrorResponse, bool) {

			logger.Error("received an error that was not handled", log.LoggerField{
				FieldName:  "err",
				FieldValue: err,
			})

			return viewmodel.HttpErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: internalServerErrorMessage,
			}, true
		}
	}
}
