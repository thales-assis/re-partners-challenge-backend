package usecase

import (
	"context"

	"github.com/re-partners-challenge-backend/internal/handler/http/viewmodel"
)

type Pack interface {
	FindAll(ctx context.Context) (viewmodel.GetPacksResponse, error)
	UpdateAll(ctx context.Context, updateAllPacks viewmodel.UpdateAllPacksRequest) error
}
