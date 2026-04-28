package httppackage

import (
	"encoding/json"
	"net/http"

	"github.com/re-partners-challenge-backend/internal/domain/contract/usecase"
	"github.com/re-partners-challenge-backend/internal/handler/http/viewmodel"
	"github.com/re-partners-challenge-backend/internal/infra/log"
	"github.com/uptrace/bunrouter"
)

type UpdatePackagesHandler bunrouter.HandlerFunc

func ProvideUpdatePackagesHandler(
	logger *log.ZapLogger,
	packageUseCase usecase.Package,
) UpdatePackagesHandler {
	return HandleUpdatePackages(logger, packageUseCase)
}

func HandleUpdatePackages(
	logger *log.ZapLogger,
	packageUseCase usecase.Package,
) UpdatePackagesHandler {
	return func(w http.ResponseWriter, req bunrouter.Request) error {

		ctx := req.Context()

		var vm viewmodel.UpdateAllPackages
		if err := json.NewDecoder(req.Body).Decode(&vm); err != nil {
			return err
		}

		if err := vm.Validate(); err != nil {
			return err
		}

		err := packageUseCase.UpdateAll(ctx, vm)
		if err != nil {
			return err
		}

		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}
