package httpserver

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(Routes), "*"),
	wire.Struct(new(ServerOption), "*"),
	ProvideCORSMiddleware,
	ProvideCoreMiddlewares,
	ProvideHandler,
	ProvideHTTPServer,
)
