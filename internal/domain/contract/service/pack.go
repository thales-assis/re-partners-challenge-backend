package service

import (
	"context"

	"github.com/re-partners-challenge-backend/internal/domain/entity"
)

type Pack interface {
	Find(ctx context.Context) ([]entity.Pack, error)
	Update(ctx context.Context, packs []entity.Pack) error
}
