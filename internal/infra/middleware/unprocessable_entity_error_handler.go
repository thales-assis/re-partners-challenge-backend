package middleware

import (
	"net/http"

	"github.com/re-partners-challenge-backend/internal/handler/http/viewmodel"
	"github.com/uptrace/bunrouter"
)

func UnprocessableEntityErrorHandler() ErrorHandlerFunc {
	return func(err error) func(w http.ResponseWriter, req bunrouter.Request) (viewmodel.HttpErrorResponse, bool) {
		return func(w http.ResponseWriter, req bunrouter.Request) (viewmodel.HttpErrorResponse, bool) {

			switch t := err.(type) {
			case *viewmodel.ValidationErrors:
				return viewmodel.HttpErrorResponse{
					Status:           http.StatusUnprocessableEntity,
					Message:          "some parameter in the request body is invalid",
					ValidationErrors: t,
				}, true
			}
			return viewmodel.HttpErrorResponse{}, false
		}
	}
}
