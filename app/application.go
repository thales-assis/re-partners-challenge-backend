package app

import (
	"github.com/re-partners-challenge-backend/internal/infra/config"
	"github.com/re-partners-challenge-backend/internal/infra/httpserver"
)

type Application struct {
	cfg    *config.Config
	server *httpserver.Server
}

func ProvideApplication(
	cfg *config.Config,
	server *httpserver.Server,
) Application {
	return Application{
		cfg,
		server,
	}
}

func (app Application) Server() *httpserver.Server {
	return app.server
}
