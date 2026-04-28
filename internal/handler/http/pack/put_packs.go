package pack

import (
	"encoding/json"
	"net/http"

	"github.com/re-partners-challenge-backend/internal/domain/contract/usecase"
	"github.com/re-partners-challenge-backend/internal/handler/http/viewmodel"
	"github.com/re-partners-challenge-backend/internal/infra/log"
	"github.com/uptrace/bunrouter"
)

type UpdatePacksHandler bunrouter.HandlerFunc

func ProvideUpdatePacksHandler(
	logger *log.ZapLogger,
	packUseCase usecase.Pack,
) UpdatePacksHandler {
	return HandleUpdatePacks(logger, packUseCase)
}

// HandleUpdatePacks godoc
// @Summary Update pack sizes
// @Description Replaces all configured package sizes.
// @Tags packs
// @Accept json
// @Produce json
// @Param payload body viewmodel.UpdateAllPacksRequest true "Pack sizes payload"
// @Success 204
// @Failure 422 {object} viewmodel.HttpErrorResponse
// @Failure 500 {object} viewmodel.HttpErrorResponse
// @Router /packs [put]
func HandleUpdatePacks(
	logger *log.ZapLogger,
	packUseCase usecase.Pack,
) UpdatePacksHandler {
	return func(w http.ResponseWriter, req bunrouter.Request) error {

		ctx := req.Context()

		var vm viewmodel.UpdateAllPacksRequest
		if err := json.NewDecoder(req.Body).Decode(&vm); err != nil {
			return err
		}

		if err := vm.Validate(); err != nil {
			return err
		}

		err := packUseCase.UpdateAll(ctx, vm)
		if err != nil {
			return err
		}

		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}
