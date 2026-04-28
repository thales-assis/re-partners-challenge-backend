package packagepersistence

import (
	"context"

	"github.com/re-partners-challenge-backend/internal/domain/contract/repository"
	"github.com/re-partners-challenge-backend/internal/domain/entity"
	"github.com/re-partners-challenge-backend/internal/infra/log"
	"github.com/re-partners-challenge-backend/internal/persistence/database"
)

type PackageRepository struct {
	logger *log.ZapLogger
	fakeDB *database.FakeDatabase
}

func ProvidePackageRepository(
	logger *log.ZapLogger,
	fakeDB *database.FakeDatabase,
) repository.Package {
	return PackageRepository{
		logger,
		fakeDB,
	}
}

func (p PackageRepository) Find(ctx context.Context) ([]entity.Package, error) {

	packages := make([]entity.Package, 0, len(p.fakeDB.Records))

	for _, pkg := range p.fakeDB.Records {
		packages = append(packages, pkg)
	}

	return packages, nil

}

func (p PackageRepository) BulkInsert(ctx context.Context, packages []entity.Package) error {

	for _, pkg := range packages {

		p.fakeDB.CountIDs++

		newRecord := entity.Package{
			ID:        p.fakeDB.CountIDs,
			CreatedAt: pkg.CreatedAt,
			Size:      pkg.Size,
		}

		p.fakeDB.Records[newRecord.ID] = newRecord
	}

	return nil
}

func (p PackageRepository) DeleteAll(ctx context.Context) error {

	p.fakeDB.Records = make(map[uint32]entity.Package)

	return nil
}
