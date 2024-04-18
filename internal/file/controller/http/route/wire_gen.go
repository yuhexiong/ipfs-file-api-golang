// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package route

import (
	"ipfs-file-api/internal/file/controller/http/handler"
)

// Injectors from wire.go:

func InitializeRouter() Router {
	fileCIDHandler := handler.InitialFileCIDHandler()
	router := NewRouter(fileCIDHandler)
	return router
}
