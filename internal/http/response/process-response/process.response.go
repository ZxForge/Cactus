package processresponse

import "cactus/internal/schema"

type ProcessesResponse struct {
	Processes []schema.Process `json:"processes"`
}