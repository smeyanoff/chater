package repository

import (
	models "chater/internal/domain/entity"
	"context"
)

type MLModelRepository interface {
	Save(ctx context.Context, model *models.MLModel) error
	Delete(ctx context.Context, modelID uint) error
	GetByID(ctx context.Context, modelID uint) (model *models.MLModel, err error)
}
