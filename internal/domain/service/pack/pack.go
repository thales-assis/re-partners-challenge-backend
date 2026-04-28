package packservice

import (
	"context"

	"github.com/re-partners-challenge-backend/internal/domain/contract/repository"
	"github.com/re-partners-challenge-backend/internal/domain/contract/service"
	"github.com/re-partners-challenge-backend/internal/domain/entity"
	"github.com/re-partners-challenge-backend/internal/infra/log"
)

type PackService struct {
	logger            *log.ZapLogger
	packageRepository repository.Pack
}

func ProvidePackService(
	logger *log.ZapLogger,
	packageRepository repository.Pack,
) service.Pack {
	return PackService{
		logger,
		packageRepository,
	}
}

func (svc PackService) Find(ctx context.Context) ([]entity.Pack, error) {

	packs, err := svc.packageRepository.Find(ctx)
	if err != nil {
		return nil, err
	}

	return packs, nil
}

func (svc PackService) Update(ctx context.Context, packs []entity.Pack) error {

	err := svc.packageRepository.DeleteAll(ctx)
	if err != nil {
		return err
	}

	err = svc.packageRepository.BulkInsert(ctx, packs)
	if err != nil {
		return err
	}

	return nil
}
