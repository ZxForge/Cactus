package core

import (
	"cactus/internal/storage/db"
	"cactus/internal/storage/file"
	"cactus/internal/storage/plugin"

	"github.com/jmoiron/sqlx"
)

type Service struct {
	db          *sqlx.DB
	storage     *db.Queries
	fileStorage *file.FileStorage
	plugins     *plugin.Storage
}

func New(db *sqlx.DB, storage *db.Queries, fileStorage *file.FileStorage, plugins *plugin.Storage) *Service {
	return &Service{
		db:          db,
		storage:     storage,
		fileStorage: fileStorage,
		plugins:     plugins,
	}
}
