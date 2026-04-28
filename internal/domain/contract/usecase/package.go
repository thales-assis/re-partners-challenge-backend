package usecase

import (
	"context"

	"github.com/re-partners-challenge-backend/internal/handler/http/viewmodel"
)

type Package interface {
	FindAll(ctx context.Context) (viewmodel.GetPackagesResponse, error)
	UpdateAll(ctx context.Context, updateAllPackages viewmodel.UpdateAllPackages) error
}
