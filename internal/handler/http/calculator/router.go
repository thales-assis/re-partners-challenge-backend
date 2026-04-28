package calculator

import (
	"github.com/uptrace/bunrouter"
)

type Router struct {
	postCalculatorPacksHandler PostCalculatorPacksHandler
}

func ProvideRouter(postCalculatorPacksHandler PostCalculatorPacksHandler) Router {
	return Router{
		postCalculatorPacksHandler,
	}
}

func (r Router) Register(router *bunrouter.Group) {
	router.POST("/calculator", bunrouter.HandlerFunc(r.postCalculatorPacksHandler))
}
