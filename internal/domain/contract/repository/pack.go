package repository

import (
	"context"

	"github.com/re-partners-challenge-backend/internal/domain/entity"
)

type Pack interface {
	Find(ctx context.Context) ([]entity.Pack, error)
	BulkInsert(ctx context.Context, packs []entity.Pack) error
	DeleteAll(ctx context.Context) error
}
