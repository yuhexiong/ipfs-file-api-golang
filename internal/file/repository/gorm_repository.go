package repository

import (
	"context"
	"fmt"
	"log/slog"

	"gorm.io/gorm"
)

type GormRepository[E any] struct {
	db *gorm.DB
}

func (r *GormRepository[E]) FindFirst(ctx context.Context, field string, value any) (*E, error) {
	var entity E
	if err := r.db.WithContext(ctx).First(&entity, fmt.Sprintf("%s = ?", field), value).Error; err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return &entity, nil
}

func (r *GormRepository[E]) Create(ctx context.Context, entity E) (*E, error) {
	if err := r.db.WithContext(ctx).Create(&entity).Error; err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return &entity, nil
}
