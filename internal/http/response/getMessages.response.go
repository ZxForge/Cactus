package response

import (
	dto "cactus/internal/DTO"
)

type GetMessagesResponse struct {
	Messages []dto.Message `json:"messages"`
}
