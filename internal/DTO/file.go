package dto

import (
	"os"

	"github.com/google/uuid"
)

type SetFile struct {
	UUID uuid.UUID
}

type GetFile struct {
	File *os.File
}
