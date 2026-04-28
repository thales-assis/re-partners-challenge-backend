package middleware

import (
	"net/http"

	"github.com/re-partners-challenge-backend/internal/handler/http/viewmodel"
	"github.com/uptrace/bunrouter"
)

type ErrorHandlerFunc func(err error) func(w http.ResponseWriter, req bunrouter.Request) (viewmodel.HttpErrorResponse, bool)
