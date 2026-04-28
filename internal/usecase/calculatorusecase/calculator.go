package calculatorusecase

import (
	"context"

	"github.com/re-partners-challenge-backend/internal/domain/contract/service"
	"github.com/re-partners-challenge-backend/internal/domain/contract/usecase"
	"github.com/re-partners-challenge-backend/internal/handler/http/viewmodel"
	"github.com/re-partners-challenge-backend/internal/infra/log"
)

type CalculatorUseCase struct {
	logger            *log.ZapLogger
	calculatorService service.Calculator
	packService       service.Pack
}

func ProvideCalculatorUseCase(
	logger *log.ZapLogger,
	calculatorService service.Calculator,
	packService service.Pack,
) usecase.Calculator {
	return CalculatorUseCase{
		logger,
		calculatorService,
		packService,
	}
}

func (u CalculatorUseCase) Calculate(ctx context.Context, amount int) ([]viewmodel.CalculatorPacksResponse, error) {

	u.logger.Info("starting to calculate the distribution of packs", log.LoggerField{
		FieldName:  "amount",
		FieldValue: amount,
	})

	packs, err := u.packService.Find(ctx)
	if err != nil {
		return nil, nil
	}

	aggregatorPacks, err := u.calculatorService.Calculate(ctx, amount, packs)
	if err != nil {
		return nil, nil
	}

	response := make([]viewmodel.CalculatorPacksResponse, 0, len(aggregatorPacks))

	for i := len(aggregatorPacks) - 1; i >= 0; i-- {
		response = append(response, viewmodel.CalculatorPacksResponse{
			PackSize: aggregatorPacks[i].PackSize,
			Quantity: aggregatorPacks[i].Quantity,
		})
	}

	u.logger.Info("all packs were calculated successfully")

	return response, nil

}
