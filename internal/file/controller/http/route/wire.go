package route

import (
	"ipfs-file-api/internal/file/controller/http/handler"

	"github.com/google/wire"
)

func InitializeRouter() Router {
	wire.Build(
		NewRouter,
		handler.InitialFileCIDHandler,
	)

	return nil
}
