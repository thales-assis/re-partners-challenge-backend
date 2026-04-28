package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bunrouter"
)

func TestRouterRegister(t *testing.T) {
	t.Parallel()

	router := bunrouter.New()
	group := router.NewGroup("")

	healthRouter := ProvideRouter(HealthHandler())
	healthRouter.Register(group)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"status":"OK"}`, rec.Body.String())
}
