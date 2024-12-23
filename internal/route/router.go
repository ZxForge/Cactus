package route

import (
	chat_service "cactus/internal/service/chat"
	"cactus/internal/service/core"
	email_service "cactus/internal/service/email"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func New(
	coreService *core.Service,
	emailService *email_service.Service,
	chatService *chat_service.Service,
) *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Подключаем модули
	addRouteApi(r, coreService, emailService)
	addRouteWS(r, chatService)
	// addRouteAuth(r)

	return r
}
