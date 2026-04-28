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
