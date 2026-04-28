package packageservice

import (
	"context"

	"github.com/re-partners-challenge-backend/internal/domain/contract/repository"
	"github.com/re-partners-challenge-backend/internal/domain/contract/service"
	"github.com/re-partners-challenge-backend/internal/domain/entity"
	"github.com/re-partners-challenge-backend/internal/infra/log"
)

type PackageService struct {
	logger            *log.ZapLogger
	packageRepository repository.Package
}

func ProvidePackageService(
	logger *log.ZapLogger,
	packageRepository repository.Package,
) service.Package {
	return PackageService{
		logger,
		packageRepository,
	}
}

func (svc PackageService) Find(ctx context.Context) ([]entity.Package, error) {

	packages, err := svc.packageRepository.Find(ctx)
	if err != nil {
		return nil, err
	}

	return packages, nil
}

func (svc PackageService) Update(ctx context.Context, packages []entity.Package) error {

	err := svc.packageRepository.DeleteAll(ctx)
	if err != nil {
		return err
	}

	err = svc.packageRepository.BulkInsert(ctx, packages)
	if err != nil {
		return err
	}

	return nil
}
