package service

import (
	"context"

	"github.com/re-partners-challenge-backend/internal/domain/entity"
)

type Calculator interface {
	Calculate(ctx context.Context, amount int, packs []entity.Pack) ([]entity.AggregatorPack, error)
}
