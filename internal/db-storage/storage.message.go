package dbstorage

import (
	"cactus/internal/model"
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/Masterminds/squirrel"
)

func (s *Storage) CreateMessage(
	ctx context.Context,
	domainId int64,
	uuid string,
	priority uint8,
	sendLater time.Time,
	value []byte,
) (model.Message, error) {
	var statusMessageID int64 = 1 // определенный статус так как при создании всегда один статутс условно 'принят'
	sql, args, _ := squirrel.
		Insert("message").
		Columns("domain_id", "status_message_id", "\"UUID\"", "value", "priority", "send_later").
		Values(domainId, statusMessageID, uuid, value, priority, sendLater).
		Suffix("RETURNING *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	var newModel = model.Message{}
	err := s.db.QueryRowxContext(ctx, sql, args...).StructScan(&newModel)
	if err != nil {
		return model.Message{}, fmt.Errorf("%s", err.Error())
	}

	return newModel, nil
}

func (s *Storage) GetStatus(ctx context.Context, uuid string) (string, error) {
	sql, args, _ := squirrel.
		Select("sm.name AS status").
		From("message AS m").
		InnerJoin("status_message AS sm ON sm.id = m.status_message_id").
		Where(squirrel.Eq{"m.\"UUID\"": uuid}).
		Limit(1).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	var status string

	err := s.db.QueryRowxContext(ctx, sql, args...).Scan(&status)
	if err != nil {
		return "", fmt.Errorf("%s", err.Error())
	}
	return status, nil
}

func (s *Storage) Abort(ctx context.Context, uuid string) (string, error) {
	var statusId int64 = 3 // Рассылка отменена
	sql, args, _ := squirrel.
		Update("message").
		Set("status_message_id", statusId).
		Where(squirrel.Eq{"\"UUID\"": uuid}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	slog.Info(sql)

	_, err := s.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return "", fmt.Errorf("%s", err.Error())
	}

	status, err := s.GetStatus(ctx, uuid)
	if err != nil {
		return "", err
	}
	return status, nil
}

func (s *Storage) GetMessagesByClienIDAndProcessID(ctx context.Context, clientId int, processId int) ([]model.Message, error) {
	sql, args, _ := squirrel.
		Select("\"UUID\"", "title", "\"value\"", "priority", "send_later", "created_at").
		From("message").
		Where(squirrel.Eq{"client_id": clientId, "process_id": processId}).
		Limit(1).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	var messages []model.Message = []model.Message{}

	rows, err := s.db.QueryxContext(ctx, sql, args...)
	if err != nil {
		return []model.Message{}, fmt.Errorf("%s", err.Error())
	}

	for rows.Next() {
		var message model.Message
		rows.StructScan(message)
		messages = append(messages, message)
	}

	return messages, nil
}
