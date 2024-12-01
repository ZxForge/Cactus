package appservice

import "cactus/internal/model"

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

func (s *Service) GetMessage() (model.Message, error) {
	return model.Message{}, nil
}
