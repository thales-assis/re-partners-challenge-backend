package pack

import (
	"net/http"

	"github.com/re-partners-challenge-backend/internal/domain/contract/usecase"
	"github.com/re-partners-challenge-backend/internal/infra/log"
	"github.com/uptrace/bunrouter"
)

type GetPacksHandler bunrouter.HandlerFunc

func ProvideGetPacksHandler(
	logger *log.ZapLogger,
	packUseCase usecase.Pack,
) GetPacksHandler {
	return HandleGetPacks(logger, packUseCase)
}

// HandleGetPacks godoc
// @Summary List pack sizes
// @Description Returns all configured package sizes.
// @Tags packs
// @Produce json
// @Success 200 {object} viewmodel.GetPacksResponse
// @Failure 500 {object} viewmodel.HttpErrorResponse
// @Router /packs [get]
func HandleGetPacks(
	logger *log.ZapLogger,
	packUseCase usecase.Pack,
) GetPacksHandler {
	return func(w http.ResponseWriter, req bunrouter.Request) error {

		ctx := req.Context()

		response, err := packUseCase.FindAll(ctx)
		if err != nil {
			return err
		}

		return bunrouter.JSON(w, response)
	}
}
