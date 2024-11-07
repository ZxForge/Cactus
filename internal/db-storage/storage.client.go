package dbstorage

import (
	"cactus/internal/model"
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func (s *Storage) GetClientByToken(ctx context.Context, token string) (model.Client, error) {

	sql, args, _ := squirrel.
		Select("id, name, token, priority").
		From("domain").
		Where(squirrel.Eq{"token": token}).
		Limit(1).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	var clients []model.Client
	errDb := s.db.SelectContext(ctx, &clients, sql, args...)
	if errDb != nil {
		return model.Client{}, fmt.Errorf("%s", errDb.Error())
	}
	if len(clients) == 0 {
		return model.Client{}, nil
	}
	return clients[0], nil
}
