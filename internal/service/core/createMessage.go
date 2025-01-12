package core

import (
	dto "cactus/internal/DTO"
	"cactus/internal/plugin"
	"cactus/internal/storage/db"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type CreateMessageParams struct {
	Plugin         plugin.Plugin
	IDSystem       int32
	PrioritySlug   string
	ChangelSlug    string
	Schema         any
	Title, Message string
	Subject        string
	SendLater      *time.Time
	Files          []SetFileParams
}

func (s *Service) CreateMessage(
	ctx context.Context,
	arg CreateMessageParams,
) (dto.CreateMessage, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("Ошибка создания контекста:", slog.String("error", err.Error()))
		return dto.CreateMessage{}, fmt.Errorf("невозможно создать транзакцию для сообщения")
	}
	defer tx.Rollback()

	storage := s.storage.WithTx(tx)

	var sendLater sql.NullTime
	if arg.SendLater == nil {
		sendLater = sql.NullTime{
			Valid: false,
		}
	} else {
		sendLater = sql.NullTime{
			Time:  *arg.SendLater,
			Valid: true,
		}
	}

	messagePriority, err := storage.GetPriorityBySlug(ctx, arg.PrioritySlug)
	if err != nil {
		return dto.CreateMessage{}, fmt.Errorf("ошибка получения приоритета по slug: %w", err)
	}

	Process, err := storage.GetTypeWorkerBySlug(ctx, arg.ChangelSlug)
	if err != nil {
		return dto.CreateMessage{}, fmt.Errorf("ошибка получения типа воркера по slug: %w", err)
	}

	// TODO: [Вынести логику redis, это не относится к созданию сообщения] после добавления redis дополнить метод удалением из очереди в redis

	value, err := json.Marshal(arg.Schema)
	if err != nil {
		return dto.CreateMessage{}, fmt.Errorf("ошибка создания json из данных: %w", err)
	}

	newMessage, err := storage.CreateMessage(
		ctx,
		db.CreateMessageParams{
			IDWorker: sql.NullInt32{
				Valid: false, // TODO: Пока NULL но надо определять по sendLater текущий Worker
			},
			IDTypeWorker: Process.ID,
			IDSystem:     arg.IDSystem,
			Uuid:         uuid.New(),
			Value:        value,
			IDPriority:   messagePriority.ID,
			SendLater:    sendLater,
		},
	)
	if err != nil {
		return dto.CreateMessage{}, fmt.Errorf("ошибка создания сообщения: %w", err)
	}

	var files []dto.SetFile
	if len(arg.Files) != 0 {
		// TODO запараллелить сохранение файлов
		for _, file := range arg.Files {
			file.IDMessage = newMessage.ID
			ff, err := s.SetFileTX(ctx, tx, file)
			if err != nil {
				return dto.CreateMessage{}, fmt.Errorf("ошибка сохранения файла: %w", err)
			}
			files = append(files, ff)
		}
	}

	dtoMessage := dto.CreateMessage{
		Message: dto.Message{
			Message: newMessage,
			Value:   arg.Schema,
		},
		Files: &files,
	}
	err = tx.Commit()
	return dtoMessage, err
}
