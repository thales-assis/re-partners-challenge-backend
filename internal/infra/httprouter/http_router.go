package httprouter

import (
	"fmt"

	handlerhttp "github.com/re-partners-challenge-backend/internal/handler/http"
	"github.com/re-partners-challenge-backend/internal/infra/config"
	"github.com/re-partners-challenge-backend/internal/infra/httpserver"

	"github.com/uptrace/bunrouter"
	"go.uber.org/zap"
)

type GroupHandler struct {
	group   *bunrouter.Group
	routers []handlerhttp.Router
}

func ProvideRouter(
	logger *zap.Logger,
	config *config.Config,
	routes httpserver.Routes,
) (*bunrouter.Router, error) {

	router := bunrouter.New()

	basePrefix := fmt.Sprintf("%s/%s", config.Server.Prefix, config.Server.Version)

	openRoutes := router.NewGroup(basePrefix)

	groupHandlers := []GroupHandler{
		{
			group:   openRoutes,
			routers: routes.Open(),
		},
	}

	for _, groupHandler := range groupHandlers {
		for _, router := range groupHandler.routers {
			router.Register(groupHandler.group)
		}
	}

	return router, nil
}
