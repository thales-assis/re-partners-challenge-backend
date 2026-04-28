package packageusecase

import (
	"context"
	"errors"
	"sort"

	"github.com/re-partners-challenge-backend/internal/domain/contract/service"
	"github.com/re-partners-challenge-backend/internal/domain/contract/usecase"
	"github.com/re-partners-challenge-backend/internal/domain/entity"
	"github.com/re-partners-challenge-backend/internal/handler/http/viewmodel"
	"github.com/re-partners-challenge-backend/internal/infra/log"
)

type PackageUseCase struct {
	logger         *log.ZapLogger
	packageService service.Package
}

func ProvidePackageUseCase(
	logger *log.ZapLogger,
	packageService service.Package,
) usecase.Package {

	return PackageUseCase{
		logger,
		packageService,
	}
}

var (
	ErrInvalidPackageSizeValue = errors.New("package size is invalid. It must be a valid value (greater than 0)")
)

func (u PackageUseCase) FindAll(ctx context.Context) (viewmodel.GetPackagesResponse, error) {

	u.logger.Info("starting to find all packages")

	packages, err := u.packageService.Find(ctx)
	if err != nil {
		return viewmodel.GetPackagesResponse{}, err
	}

	response := viewmodel.GetPackagesResponse{
		Sizes: make([]int, 0, len(packages)),
	}

	for _, packge := range packages {

		response.Sizes = append(response.Sizes, int(packge.Size))
	}

	sort.Ints(response.Sizes)

	return response, nil
}

func (u PackageUseCase) UpdateAll(ctx context.Context, updateAllPackages viewmodel.UpdateAllPackages) error {

	u.logger.Info("starting to update all packages")

	newPackages := make([]entity.Package, 0, len(updateAllPackages.Sizes))

	for _, newPackageSize := range updateAllPackages.Sizes {

		if newPackageSize <= 0 {
			return ErrInvalidPackageSizeValue
		} else {
			newPackage := entity.NewPackage(newPackageSize)
			newPackages = append(newPackages, newPackage)
		}
	}

	err := u.packageService.Update(ctx, newPackages)
	if err != nil {
		return err
	}

	u.logger.Info("all packages were updated successfully")

	return nil
}
