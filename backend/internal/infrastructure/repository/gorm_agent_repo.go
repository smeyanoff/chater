package repository

import (
	entities "chater/internal/domain/entity"
	"chater/internal/domain/repository"
	"context"

	"gorm.io/gorm"
)

type gormAgentRepo struct {
	db *gorm.DB
}

func NewGormAgentRepo(db *gorm.DB) repository.AgentRepository {
	return &gormAgentRepo{db: db}
}

func (r *gormAgentRepo) Save(ctx context.Context, agent *entities.Agent) error {
	return r.db.WithContext(ctx).Create(agent).Error
}

func (r *gormAgentRepo) Delete(ctx context.Context, agentID uint) error {
	return r.db.WithContext(ctx).Delete(&entities.Agent{}, agentID).Error
}

func (r *gormAgentRepo) AddRoute(ctx context.Context, route *entities.AgentRoute) error {
	return r.db.WithContext(ctx).Create(route).Error
}
