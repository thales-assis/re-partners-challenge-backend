package httpserver

import (
	"net/http"

	"github.com/rs/cors"
	"github.com/uptrace/bunrouter"
)

func ProvideHandler(
	cors *cors.Cors,
	router *bunrouter.Router,
) http.Handler {

	handler := cors.Handler(router)

	return handler
}
