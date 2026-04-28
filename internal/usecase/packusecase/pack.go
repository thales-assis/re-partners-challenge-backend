package packusecase

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

type PackUseCase struct {
	logger      *log.ZapLogger
	packService service.Pack
}

func ProvidePackUseCase(
	logger *log.ZapLogger,
	packService service.Pack,
) usecase.Pack {

	return PackUseCase{
		logger,
		packService,
	}
}

var (
	ErrInvalidPackSizeValue = errors.New("package size is invalid. It must be a valid value (greater than 0)")
)

func (u PackUseCase) FindAll(ctx context.Context) (viewmodel.GetPacksResponse, error) {

	u.logger.Info("starting to find all packs")

	packs, err := u.packService.Find(ctx)
	if err != nil {
		return viewmodel.GetPacksResponse{}, err
	}

	response := viewmodel.GetPacksResponse{
		Sizes: make([]int, 0, len(packs)),
	}

	for _, packge := range packs {

		response.Sizes = append(response.Sizes, int(packge.Size))
	}

	sort.Ints(response.Sizes)

	return response, nil
}

func (u PackUseCase) UpdateAll(ctx context.Context, updateAllPacks viewmodel.UpdateAllPacksRequest) error {

	u.logger.Info("starting to update all packs", log.LoggerField{
		FieldName:  "packs",
		FieldValue: updateAllPacks.Sizes,
	})

	newPacks := make([]entity.Pack, 0, len(updateAllPacks.Sizes))

	for _, newPackSize := range updateAllPacks.Sizes {

		if newPackSize <= 0 {
			return ErrInvalidPackSizeValue
		} else {
			newPack := entity.NewPack(newPackSize)
			newPacks = append(newPacks, newPack)
		}
	}

	err := u.packService.Update(ctx, newPacks)
	if err != nil {
		return err
	}

	u.logger.Info("all packs were updated successfully")

	return nil
}
