package emailrequest

import "time"

type SendRequest struct {
	Title          string             `schema:"title" validate:"required,max=255"`
	Type           *string            `schema:"type" validate:"omitnil,oneof=text html template"`
	Message        string             `schema:"message" validate:"required"`
	Subjects       []string           `schema:"subjects[]" validate:"required,min=1"`
	Data           *map[string]string `schema:"data[]" validate:"required_if=Type template,omitempty,min=1"`
	Priority       *uint8             `schema:"priority" validate:"required,min=0,max=5"`
	SendLater      *time.Time         `schema:"send_later" validate:"omitnil,required"` // Сомнительно но окей
	SameAttachment *bool              `schema:"same_attachment" validate:"required,omitnil,boolean"`
	Theme          *string            `schema:"theme" validate:"required"`
	Headers        *map[string]string `schema:"headers[]" validate:"omitnil"`
	Files          []*struct {
		Title *string `schema:"title" validate:"required,min=5,max=255"`
		Name  *string `schema:"name" validate:"required"`
		Type  *string `schema:"type" validate:"required,oneof=bin base64"`
	} `schema:"files[]" validate:"omitempty,dive"`
}
