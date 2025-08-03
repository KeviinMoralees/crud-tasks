package gateways

import (
	"context"
	"crud-tasks/domain/models"
)

type TaskGateway interface {
	Create(ctx context.Context, task models.Task) (error, models.Task)
	GetById(ctx context.Context, id string) (error, models.Task)
	Delete(ctx context.Context, id string) error
}
