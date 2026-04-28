package pack

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/re-partners-challenge-backend/internal/domain/entity"
	packservice "github.com/re-partners-challenge-backend/internal/domain/service/pack"
	"github.com/re-partners-challenge-backend/internal/infra/log"
	"github.com/re-partners-challenge-backend/internal/persistence/database"
	"github.com/re-partners-challenge-backend/internal/persistence/packpersistence"
	"github.com/re-partners-challenge-backend/internal/usecase/packusecase"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bunrouter"
)

func provideFakeDB() *database.FakeDatabase {

	fakeDB := &database.FakeDatabase{
		CountIDs: 3,
		Records: map[uint32]entity.Pack{
			1: {
				ID:        1,
				CreatedAt: time.Date(2026, 4, 28, 18, 30, 0, 0, time.UTC),
				Size:      53,
			},
			2: {
				ID:        2,
				CreatedAt: time.Date(2026, 4, 28, 18, 30, 0, 0, time.UTC),
				Size:      23,
			},
			3: {
				ID:        3,
				CreatedAt: time.Date(2026, 4, 28, 18, 30, 0, 0, time.UTC),
				Size:      31,
			},
		},
	}

	return fakeDB
}

func TestRouterRegisterGetPacksSuccess(t *testing.T) {
	t.Parallel()

	logger, _ := log.ProvideLogger()

	fakeDB := provideFakeDB()

	packRepository := packpersistence.ProvidePackRepository(logger, fakeDB)
	packService := packservice.ProvidePackService(logger, packRepository)
	packUseCase := packusecase.ProvidePackUseCase(logger, packService)

	getPacksHandler := ProvideGetPacksHandler(logger, packUseCase)

	router := bunrouter.New()
	group := router.NewGroup("")

	packRouter := ProvideRouter(getPacksHandler, nil)
	packRouter.Register(group)

	req := httptest.NewRequest(http.MethodGet, "/packs", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"pack_sizes":[23,31,53]}`, rec.Body.String())
}

func TestRouterRegisterPutPacksSuccess(t *testing.T) {
	t.Parallel()

	logger, _ := log.ProvideLogger()

	fakeDB := provideFakeDB()

	packRepository := packpersistence.ProvidePackRepository(logger, fakeDB)
	packService := packservice.ProvidePackService(logger, packRepository)
	packUseCase := packusecase.ProvidePackUseCase(logger, packService)

	getPacksHandler := ProvideGetPacksHandler(logger, packUseCase)
	updatePacksHandler := ProvideUpdatePacksHandler(logger, packUseCase)

	router := bunrouter.New()
	group := router.NewGroup("")

	packRouter := ProvideRouter(getPacksHandler, updatePacksHandler)
	packRouter.Register(group)

	req := httptest.NewRequest(http.MethodGet, "/packs", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"pack_sizes":[23,31,53]}`, rec.Body.String())

	body := `{"pack_sizes":[250, 500, 1000, 2000, 5000]}`
	req = httptest.NewRequest(http.MethodPut, "/packs", strings.NewReader(body))
	rec = httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	req = httptest.NewRequest(http.MethodGet, "/packs", nil)
	rec = httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"pack_sizes":[250, 500, 1000, 2000, 5000]}`, rec.Body.String())
}
