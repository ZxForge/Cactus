package route

import (
	chat_service "cactus/internal/service/chat"
	email_service "cactus/internal/service/email"

	"github.com/go-chi/chi/v5"
)

func New(
	emailService *email_service.Service,
	chatService *chat_service.Service,
) *chi.Mux {
	r := chi.NewRouter()

	// Подключаем модули
	addRouteApi(r, emailService)
	addRouteWS(r, chatService)
	// addRouteAuth(r)

	return r
}
