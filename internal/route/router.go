package route

import (
	chat_service "cactus/internal/service/chat"
	"cactus/internal/service/core"
	email_service "cactus/internal/service/email"

	"github.com/go-chi/chi/v5"
)

func New(
	coreService *core.Service,
	emailService *email_service.Service,
	chatService *chat_service.Service,
) *chi.Mux {
	r := chi.NewRouter()

	// Подключаем модули
	addRouteApi(r, coreService, emailService)
	addRouteWS(r, chatService)
	// addRouteAuth(r)

	return r
}
