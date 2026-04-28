package service

import (
	"context"

	"github.com/re-partners-challenge-backend/internal/domain/entity"
)

type Package interface {
	Find(ctx context.Context) ([]entity.Package, error)
	Update(ctx context.Context, packages []entity.Package) error
}
