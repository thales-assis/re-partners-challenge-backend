package httprouter

import (
	"fmt"

	_ "github.com/re-partners-challenge-backend/docs"
	handlerhttp "github.com/re-partners-challenge-backend/internal/handler/http"
	"github.com/re-partners-challenge-backend/internal/infra/config"
	"github.com/re-partners-challenge-backend/internal/infra/httpserver"
	"github.com/re-partners-challenge-backend/internal/infra/log"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/uptrace/bunrouter"
)

type GroupHandler struct {
	group   *bunrouter.Group
	routers []handlerhttp.Router
}

func ProvideRouter(
	logger *log.ZapLogger,
	config *config.Config,
	routes httpserver.Routes,
	middlewaresFunc ...bunrouter.MiddlewareFunc,
) (*bunrouter.Router, error) {

	middlewares := bunrouter.Use(middlewaresFunc...)

	router := bunrouter.New(middlewares)

	basePrefix := fmt.Sprintf("%s/%s", config.Server.Prefix, config.Server.Version)

	openRoutes := router.NewGroup(basePrefix)

	apiRoutes := openRoutes.Use( /* add some middlewares (for exemple: auth) */ )

	groupHandlers := []GroupHandler{
		{
			group:   openRoutes,
			routers: routes.Open(),
		},
		{
			group:   apiRoutes,
			routers: routes.API(),
		},
	}

	for _, groupHandler := range groupHandlers {
		for _, router := range groupHandler.routers {
			router.Register(groupHandler.group)
		}
	}

	router.GET(
		fmt.Sprintf("%s/swagger/*path", basePrefix),
		bunrouter.HTTPHandler(
			httpSwagger.Handler(
				httpSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", basePrefix)),
			),
		),
	)

	return router, nil
}
