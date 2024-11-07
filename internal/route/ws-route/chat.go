package wsroute

import (
	"cactus/internal/http/controllers/chat"
	chatservice "cactus/internal/service/chat-service"

	"github.com/go-chi/chi/v5"
)

func New(
	chatService *chatservice.Service,
) *chi.Mux {
	r := chi.NewRouter()
	// TODO: вынести выше, тут запуск сервера WS
	go chatService.Listen()

	// r.Use(middleware.CheckDomainToken(emailService))

	r.Handle("/chat/{chatID}", chat.Chat(chatService))

	return r
}
