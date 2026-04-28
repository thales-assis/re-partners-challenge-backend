package calculator

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/re-partners-challenge-backend/internal/domain/entity"
	calculatorservice "github.com/re-partners-challenge-backend/internal/domain/service/calculator"
	packservice "github.com/re-partners-challenge-backend/internal/domain/service/pack"
	"github.com/re-partners-challenge-backend/internal/infra/log"
	"github.com/re-partners-challenge-backend/internal/persistence/database"
	"github.com/re-partners-challenge-backend/internal/persistence/packpersistence"
	"github.com/re-partners-challenge-backend/internal/usecase/calculatorusecase"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bunrouter"
)

func TestRouterRegisterPostCalculatorSuccess(t *testing.T) {
	t.Parallel()

	logger, _ := log.ProvideLogger()

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

	packRepository := packpersistence.ProvidePackRepository(logger, fakeDB)
	packService := packservice.ProvidePackService(logger, packRepository)
	calculatorService := calculatorservice.ProvideCalculatorService(logger)
	calculatorUseCase := calculatorusecase.ProvideCalculatorUseCase(logger, calculatorService, packService)

	getPacksHandler := ProvidePostCalculatorPackHandler(logger, calculatorUseCase)

	router := bunrouter.New()
	group := router.NewGroup("")

	packRouter := ProvideRouter(getPacksHandler)
	packRouter.Register(group)

	body := `
				{
					"amount":500000
				}
			`

	req := httptest.NewRequest(http.MethodPost, "/calculator", strings.NewReader(body))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t,
		`[
			{
				"pack_size": 23,
				"quantity": 2
			},
			{
				"pack_size": 31,
				"quantity": 7
			},
			{
				"pack_size": 53,
				"quantity": 9429
			}
		]`,
		rec.Body.String(),
	)

}

func TestRouterRegisterPostCalculatorSuccess_2(t *testing.T) {
	t.Parallel()

	logger, _ := log.ProvideLogger()

	fakeDB := &database.FakeDatabase{
		CountIDs: 3,
		Records: map[uint32]entity.Pack{
			1: {
				ID:        1,
				CreatedAt: time.Date(2026, 4, 28, 18, 30, 0, 0, time.UTC),
				Size:      250,
			},
			2: {
				ID:        2,
				CreatedAt: time.Date(2026, 4, 28, 18, 30, 0, 0, time.UTC),
				Size:      500,
			},
			3: {
				ID:        3,
				CreatedAt: time.Date(2026, 4, 28, 18, 30, 0, 0, time.UTC),
				Size:      1000,
			},
			4: {
				ID:        4,
				CreatedAt: time.Date(2026, 4, 28, 18, 30, 0, 0, time.UTC),
				Size:      2000,
			},
			5: {
				ID:        5,
				CreatedAt: time.Date(2026, 4, 28, 18, 30, 0, 0, time.UTC),
				Size:      5000,
			},
		},
	}

	packRepository := packpersistence.ProvidePackRepository(logger, fakeDB)
	packService := packservice.ProvidePackService(logger, packRepository)
	calculatorService := calculatorservice.ProvideCalculatorService(logger)
	calculatorUseCase := calculatorusecase.ProvideCalculatorUseCase(logger, calculatorService, packService)

	getPacksHandler := ProvidePostCalculatorPackHandler(logger, calculatorUseCase)

	router := bunrouter.New()
	group := router.NewGroup("")

	packRouter := ProvideRouter(getPacksHandler)
	packRouter.Register(group)

	body := `
				{
					"amount":12001
				}
			`

	req := httptest.NewRequest(http.MethodPost, "/calculator", strings.NewReader(body))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(
		t,
		`[
			{
				"pack_size": 250,
				"quantity": 1
			},
			{
				"pack_size": 2000,
				"quantity": 1
			},
			{
				"pack_size": 5000,
				"quantity": 2
			}
		]`,
		rec.Body.String(),
	)
}
