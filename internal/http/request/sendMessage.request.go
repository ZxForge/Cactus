package request

import (
	"time"
)

type SendMessageRequestFile struct {
	Title string `json:"title" validate:"required,min=5,max=255"`
	Form  string `json:"form" validate:"required,ascii"`
}

type SendMessageRequest struct {
	Title        string                   `json:"title" validate:"required,max=255"`
	Message      string                   `json:"message" validate:"required"`
	PrioritySlug string                   `json:"priority_slug" validate:"required,min=0,max=5"`
	Subject      string                   `json:"subject" validate:"required,min=1"`
	SendLater    *time.Time               `json:"send_later" validate:"omitnil,required"` // Сомнительно но окей
	Value        string                   `json:"value" validate:"required,json"`
	Files        []SendMessageRequestFile `json:"files" validate:"omitempty,dive"`
}
