package middleware

import (
	"cactus/internal/model"
	"context"
	"net/http"
)

type EmailService interface {
	GetClientByToken(ctx context.Context, token string) (model.Client, error)
}

func CheckDomainToken(s EmailService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("X-Token-Domain") // TODO в env
			ctxReq := r.Context()
			client, err := s.GetClientByToken(ctxReq, token)
			if err != nil {
				// TODO вывести ошибку или бросить выше
			}
			if client == (model.Client{}) {
				rw.Header().Set("Content-Type", "application/json")
				rw.WriteHeader(401)
				return
			}

			ctx := context.WithValue(ctxReq, "domain-id", client.Id)
			ctx2 := context.WithValue(ctx, "domain-priority", client.Priority)
			next.ServeHTTP(rw, r.WithContext(ctx2))
		})
	}
}
