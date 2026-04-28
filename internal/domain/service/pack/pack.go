package packservice

import (
	"context"

	"github.com/re-partners-challenge-backend/internal/domain/contract/repository"
	"github.com/re-partners-challenge-backend/internal/domain/contract/service"
	"github.com/re-partners-challenge-backend/internal/domain/entity"
	"github.com/re-partners-challenge-backend/internal/infra/log"
)

type PackService struct {
	logger         *log.ZapLogger
	packRepository repository.Pack
}

func ProvidePackService(
	logger *log.ZapLogger,
	packRepository repository.Pack,
) service.Pack {
	return PackService{
		logger,
		packRepository,
	}
}

func (svc PackService) Find(ctx context.Context) ([]entity.Pack, error) {

	packs, err := svc.packRepository.Find(ctx)
	if err != nil {
		return nil, err
	}

	return packs, nil
}

func (svc PackService) Update(ctx context.Context, packs []entity.Pack) error {

	err := svc.packRepository.DeleteAll(ctx)
	if err != nil {
		return err
	}

	err = svc.packRepository.BulkInsert(ctx, packs)
	if err != nil {
		return err
	}

	return nil
}
