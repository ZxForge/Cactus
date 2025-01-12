package request

type GetFilesRequest struct {
	Hashes []string `json:"hashes" validate:"required,gt=0,dive,required"`
}
