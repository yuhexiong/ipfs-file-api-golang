package service

import (
	"ipfs-file-api/internal/file/entity"
	"ipfs-file-api/internal/file/repository"

	"github.com/google/wire"
)

var (
	FileCIDServiceSet = wire.NewSet(NewFileCIDService, repository.InitialFileCIDRepository)
)

func InitialFileCIDService() entity.FileCIDService {
	wire.Build(FileCIDServiceSet)

	return nil
}
