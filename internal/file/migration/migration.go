package migration

import (
	"context"
	"ipfs-file-api/internal/file/database"
	"ipfs-file-api/internal/file/entity"
)

func Init(ctx context.Context) (err error) {
	db := database.GetDB().WithContext(ctx)

	return db.AutoMigrate(
		&entity.FileCID{},
	)
}
