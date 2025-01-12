package request

type GetMessageForRequest struct {
	ClientId     int `json:"client_id" validate:"required,min=1"`
	TypeWorkerId int `json:"type_worker_id" validate:"required,min=1"`
}
