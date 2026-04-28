package packpersistence

import (
	"context"

	"github.com/re-partners-challenge-backend/internal/domain/contract/repository"
	"github.com/re-partners-challenge-backend/internal/domain/entity"
	"github.com/re-partners-challenge-backend/internal/infra/log"
	"github.com/re-partners-challenge-backend/internal/persistence/database"
)

type PackRepository struct {
	logger *log.ZapLogger
	fakeDB *database.FakeDatabase
}

func ProvidePackRepository(
	logger *log.ZapLogger,
	fakeDB *database.FakeDatabase,
) repository.Pack {
	return PackRepository{
		logger,
		fakeDB,
	}
}

func (p PackRepository) Find(ctx context.Context) ([]entity.Pack, error) {

	packs := make([]entity.Pack, 0, len(p.fakeDB.Records))

	for _, pkg := range p.fakeDB.Records {
		packs = append(packs, pkg)
	}

	return packs, nil

}

func (p PackRepository) BulkInsert(ctx context.Context, packs []entity.Pack) error {

	for _, pkg := range packs {

		p.fakeDB.CountIDs++

		newRecord := entity.Pack{
			ID:        p.fakeDB.CountIDs,
			CreatedAt: pkg.CreatedAt,
			Size:      pkg.Size,
		}

		p.fakeDB.Records[newRecord.ID] = newRecord
	}

	return nil
}

func (p PackRepository) DeleteAll(ctx context.Context) error {

	p.fakeDB.Records = make(map[uint32]entity.Pack)

	return nil
}
