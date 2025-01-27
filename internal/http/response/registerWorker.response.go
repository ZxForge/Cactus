package response

type RegisterWorkerResponse struct {
	Created bool  `json:"created"`
	Id      int32 `json:"id"`
}
