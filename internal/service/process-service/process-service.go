package processservice

import (
	"cactus/internal/model"
	"cactus/internal/schema"
	"context"
	"fmt"
)

type Storage interface {
	GetProcesses(ctx context.Context) ([]model.Process, error)
}

type Service struct {
	storage Storage
}

func New(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) GetProcesses(ctx context.Context) ([]schema.Process, error) {
	processmodel, err := s.storage.GetProcesses(ctx)
	if err != nil {
		return []schema.Process{}, fmt.Errorf("%v", err.Error())
	}

	var processes []schema.Process = []schema.Process{}

	for _, p := range processmodel {
		var process schema.Process
		process.Id = p.Id
		process.Name = p.Name
		process.Slug = p.Slug
		processes = append(processes, process)
	}

	return processes, err 
}