package repository

import (
	"ipfs-file-api/internal/file/entity"

	"gorm.io/gorm"
)

type fileCIDRepository struct {
	*GormRepository[entity.FileCID]
}

func NewFileCIDRepository(db *gorm.DB) entity.FileCIDRepository {
	return &fileCIDRepository{
		GormRepository: &GormRepository[entity.FileCID]{db},
	}
}
