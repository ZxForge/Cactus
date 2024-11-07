package dbstorage

import (
	"cactus/internal/model"
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func (s *Storage) GetUserById(ctx context.Context, id int64) (model.User, error) {
	sql, args, _ := squirrel.
		Select("id, role_id, first_name, last_name, email, password, create_at").
		From("users").
		Where(squirrel.Eq{"id": id}).
		Limit(1).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	fmt.Printf("%v\n", sql)

	var user []model.User

	errDb := s.db.SelectContext(ctx, &user, sql, args...)
	if errDb != nil {
		return model.User{}, fmt.Errorf("%s", errDb.Error())
	}
	return user[0], nil
}
