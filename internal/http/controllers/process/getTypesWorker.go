package process

import (
	dto "cactus/internal/DTO"
	type_worker_response "cactus/internal/http/response/type_worker"
	"context"
	"encoding/json"
	"net/http"
)

type getTypesWorkerService interface {
	GetTypeWorkers(ctx context.Context) ([]dto.TypeWorker, error)
}

func GetTypesWorker(s getTypesWorkerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO обработать ошибку
		processes, _ := s.GetTypeWorkers(r.Context())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		jsonResp, _ := json.Marshal(type_worker_response.CreateResponseOK(type_worker_response
			: processes,
		}))
		w.Write(jsonResp)
	}
}
