package core

import (
	"cactus/internal/storage/db"
	"context"
)

func (s *Service) GetMessage(ctx context.Context) (db.Message, error) {
	return db.Message{}, nil
}
