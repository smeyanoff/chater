package repository

import (
	models "chater/internal/domain/entity"
	"context"
)

type AgentRepository interface {
	Save(ctx context.Context, agent *models.Agent) error
	Delete(ctx context.Context, agentID uint) error
	AddRoute(ctx context.Context, route *models.AgentRoute) error
}
