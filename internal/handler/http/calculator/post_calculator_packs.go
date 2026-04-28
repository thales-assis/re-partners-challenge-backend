package calculator

import (
	"encoding/json"
	"net/http"

	"github.com/re-partners-challenge-backend/internal/domain/contract/usecase"
	"github.com/re-partners-challenge-backend/internal/handler/http/viewmodel"
	"github.com/re-partners-challenge-backend/internal/infra/log"
	"github.com/uptrace/bunrouter"
)

type PostCalculatorPacksHandler bunrouter.HandlerFunc

func ProvidePostCalculatorPackHandler(
	logger *log.ZapLogger,
	calculatorUseCase usecase.Calculator,
) PostCalculatorPacksHandler {
	return HandleCalculatorPack(logger, calculatorUseCase)
}

func HandleCalculatorPack(
	logger *log.ZapLogger,
	calculatorUseCase usecase.Calculator,
) PostCalculatorPacksHandler {
	return func(w http.ResponseWriter, req bunrouter.Request) error {

		ctx := req.Context()

		var vm viewmodel.CalculatorPacksRequest
		if err := json.NewDecoder(req.Body).Decode(&vm); err != nil {
			return err
		}

		if err := vm.Validate(); err != nil {
			return err
		}

		response, err := calculatorUseCase.Calculate(ctx, vm.Amount)
		if err != nil {
			return err
		}

		return bunrouter.JSON(w, response)
	}
}
