package postgres

import (
	"context"

	"github.com/viacheslavek/grpcauth/auth/internal/domain/models"
)

// TODO: нужно в тестах -> техдолг

func (s *Storage) GetApp(ctx context.Context, appId int) (models.App, error) {
	//TODO implement me
	panic("implement me")
}