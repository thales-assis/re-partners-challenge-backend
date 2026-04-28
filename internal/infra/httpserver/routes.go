package httpserver

import (
	"github.com/re-partners-challenge-backend/internal/handler/http"
	"github.com/re-partners-challenge-backend/internal/handler/http/calculator"
	"github.com/re-partners-challenge-backend/internal/handler/http/health"
	"github.com/re-partners-challenge-backend/internal/handler/http/pack"
)

type Routes struct {
	HealthCheckRouter health.Router
	PackRouter        pack.Router
	CalculatorRouter  calculator.Router
}

func (r Routes) Open() []http.Router {
	return []http.Router{
		r.HealthCheckRouter,
	}
}

func (r Routes) API() []http.Router {
	return []http.Router{
		r.PackRouter,
		r.CalculatorRouter,
	}
}
