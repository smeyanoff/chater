package repository

import (
	entities "chater/internal/domain/entity"
	"chater/internal/domain/repository"
	"context"

	"gorm.io/gorm"
)

type gormMLAgentRepo struct {
	db *gorm.DB
}

func NewGormAgentRepo(db *gorm.DB) repository.MLAgentRepository {
	return &gormMLAgentRepo{db: db}
}

func (r *gormMLAgentRepo) Save(ctx context.Context, agent *entities.Agent) error {
	return r.db.WithContext(ctx).Create(agent).Error
}

func (r *gormMLAgentRepo) Delete(ctx context.Context, agentID uint) error {
	return r.db.WithContext(ctx).Delete(&entities.Agent{}, agentID).Error
}

func (r *gormMLAgentRepo) GetByID(ctx context.Context, agentID uint) (agent *entities.Agent, err error) {
	if err := r.db.WithContext(ctx).First(agent, agentID).Error; err != nil {
		return nil, err
	}
	return agent, nil
}
