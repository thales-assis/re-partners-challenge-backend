package httpserver

import (
	"github.com/re-partners-challenge-backend/internal/infra/config"

	"github.com/rs/cors"
)

func ProvideCORSMiddleware(c *config.Config) *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: c.Server.Cors.AllowedOrigins,
		AllowedMethods: c.Server.Cors.AllowedMethods,
		AllowedHeaders: c.Server.Cors.AllowedHeaders,
	})
}
