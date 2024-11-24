package processController

import (
	processresponse "cactus/internal/http/response/process-response"
	"cactus/internal/schema"
	"context"
	"encoding/json"
	"net/http"
) 

type processService interface {
	GetProcesses(ctx context.Context) ([]schema.Process, error)
}

func GetProcessesFor(s processService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO обработать ошибку
		processes,  _:= s.GetProcesses(r.Context())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		jsonResp, _ := json.Marshal(processresponse.CreateResponseOK(processresponse.ProcessesResponse{
			Processes: processes,
		}))
		w.Write(jsonResp)
	}
}
