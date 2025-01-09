package repository

import (
	models "chater/internal/domain/entity"
	"context"
)

type MLAgentRepository interface {
	Save(ctx context.Context, agent *models.Agent) error
	Delete(ctx context.Context, agentID uint) error
	GetByID(ctx context.Context, agentID uint) (agent *models.Agent, err error)
}
