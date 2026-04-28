package http

import (
	"github.com/uptrace/bunrouter"
)

type Router interface {
	Register(router *bunrouter.Group)
}
