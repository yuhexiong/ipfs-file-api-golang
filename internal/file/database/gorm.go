package database

import (
	"ipfs-file-api/internal/config"
	pkgDatabase "ipfs-file-api/pkg/database"
	"sync"

	"gorm.io/gorm"
)

var (
	once sync.Once
	cfg  *pkgDatabase.Config
)

func GetDB() *gorm.DB {
	once.Do(func() {
		cfg = &pkgDatabase.Config{
			Host:     config.PostgresHost,
			Port:     config.PostgresPort,
			User:     config.PostgresUser,
			Password: config.PostgresPassword,
			DataBase: config.PostgresDataBase,
		}
	})

	return pkgDatabase.GetDB(cfg)
}
