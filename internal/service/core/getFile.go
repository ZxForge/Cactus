package core

import (
	dto "cactus/internal/DTO"
	"context"
	"fmt"

	u "github.com/google/uuid"
)

func (s *Service) GetFile(ctx context.Context, uuid string) (dto.GetFile, error) {

	UUID, _ := u.Parse(uuid)
	_ = UUID
	// TODO получить файл по uuid
	path := "go.jpeg"

	file, err := s.fileStorage.Get(ctx, path)
	if err != nil {
		return dto.GetFile{}, fmt.Errorf("%v", err.Error())
	}

	return dto.GetFile{
		File: file,
	}, nil
}
