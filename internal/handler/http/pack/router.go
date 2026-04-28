package pack

import (
	"github.com/uptrace/bunrouter"
)

type Router struct {
	getPacksHandler    GetPacksHandler
	updatePacksHandler UpdatePacksHandler
}

func ProvideRouter(
	getPacksHandler GetPacksHandler,
	updatePacksHandler UpdatePacksHandler,
) Router {
	return Router{
		getPacksHandler,
		updatePacksHandler,
	}
}

func (r Router) Register(router *bunrouter.Group) {
	router.GET("/packs", bunrouter.HandlerFunc(r.getPacksHandler))
	router.PUT("/packs", bunrouter.HandlerFunc(r.updatePacksHandler))
}
