package repository

import (
	"context"

	"github.com/re-partners-challenge-backend/internal/domain/entity"
)

type Package interface {
	Find(ctx context.Context) ([]entity.Package, error)
	BulkInsert(ctx context.Context, packages []entity.Package) error
	DeleteAll(ctx context.Context) error
}
