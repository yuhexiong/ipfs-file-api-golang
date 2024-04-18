//go:build wireinject
// +build wireinject

package repository

import (
	"ipfs-file-api/internal/file/database"
	"ipfs-file-api/internal/file/entity"

	"github.com/google/wire"
)

var (
	FileCIDRepositorySet = wire.NewSet(NewFileCIDRepository, database.GetDB)
)

func InitialFileCIDRepository() entity.FileCIDRepository {
	wire.Build(FileCIDRepositorySet)

	return nil
}
