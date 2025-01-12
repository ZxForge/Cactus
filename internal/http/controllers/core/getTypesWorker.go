package core

import (
	dto "cactus/internal/DTO"
	"cactus/internal/http/response"
	"context"
	"net/http"
)

type getTypesWorkerService interface {
	GetTypeWorkers(ctx context.Context) ([]dto.TypeWorker, error)
}

func GetTypesWorker(s getTypesWorkerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO обработать ошибку
		type_worker, _ := s.GetTypeWorkers(r.Context())
		response.ResponseOKJSON(w, response.GetTypesWorkerResponse{
			TypesWorker: type_worker,
		})
	}
}
