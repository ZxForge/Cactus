package emailrequest

type GetStatusRequest struct {
	UUID string `json:"uuid" validate:"required,uuid4"`
}
