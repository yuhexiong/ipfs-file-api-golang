//go:build wireinject
// +build wireinject

package handler

import (
	"ipfs-file-api/internal/file/service"

	"github.com/google/wire"
)

var (
	FileCIDHandlerSet = wire.NewSet(NewFileCIDHandler, service.InitialFileCIDService)
)

func InitialFileCIDHandler() *FileCIDHandler {
	wire.Build(FileCIDHandlerSet)

	return nil
}
