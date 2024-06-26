// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package repository

import (
	"github.com/google/wire"
	"ipfs-file-api/internal/file/database"
	"ipfs-file-api/internal/file/entity"
)

// Injectors from wire.go:

func InitialFileCIDRepository() entity.FileCIDRepository {
	db := database.GetDB()
	entityFileCIDRepository := NewFileCIDRepository(db)
	return entityFileCIDRepository
}

// wire.go:

var (
	FileCIDRepositorySet = wire.NewSet(NewFileCIDRepository, database.GetDB)
)
