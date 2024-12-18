package appservice

import "cactus/internal/storage/db"

type Storage interface {
}

type Service struct {
	storage Storage
}

func New(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) GetMessage() (db.Message, error) {
	return db.Message{}, nil
}
