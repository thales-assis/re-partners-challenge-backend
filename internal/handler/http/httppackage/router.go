package httppackage

import (
	"github.com/uptrace/bunrouter"
)

type Router struct {
	getPackagesHandler    GetPackagesHandler
	updatePackagesHandler UpdatePackagesHandler
}

func ProvideRouter(
	getPackagesHandler GetPackagesHandler,
	updatePackagesHandler UpdatePackagesHandler,
) Router {
	return Router{
		getPackagesHandler,
		updatePackagesHandler,
	}
}

func (r Router) Register(router *bunrouter.Group) {
	router.GET("/packages", bunrouter.HandlerFunc(r.getPackagesHandler))
	router.PUT("/packages", bunrouter.HandlerFunc(r.updatePackagesHandler))
}
