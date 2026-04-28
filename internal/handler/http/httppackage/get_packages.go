package httppackage

import (
	"net/http"

	"github.com/re-partners-challenge-backend/internal/domain/contract/usecase"
	"github.com/re-partners-challenge-backend/internal/infra/log"
	"github.com/uptrace/bunrouter"
)

type GetPackagesHandler bunrouter.HandlerFunc

func ProvideGetPackagesHandler(
	logger *log.ZapLogger,
	packageUseCase usecase.Package,
) GetPackagesHandler {
	return HandleGetPackages(logger, packageUseCase)
}

func HandleGetPackages(
	logger *log.ZapLogger,
	packageUseCase usecase.Package,
) GetPackagesHandler {
	return func(w http.ResponseWriter, req bunrouter.Request) error {

		ctx := req.Context()

		response, err := packageUseCase.FindAll(ctx)
		if err != nil {
			return err
		}

		return bunrouter.JSON(w, response)
	}
}
