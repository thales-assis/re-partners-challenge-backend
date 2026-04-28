package usecase

import (
	"context"

	"github.com/re-partners-challenge-backend/internal/handler/http/viewmodel"
)

type Calculator interface {
	Calculate(ctx context.Context, amount int) ([]viewmodel.CalculatorPacksResponse, error)
}
