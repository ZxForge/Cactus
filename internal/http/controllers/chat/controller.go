package chat

import (
	chatservice "cactus/internal/service/chat-service"
	"log/slog"

	"net/http"
)

// Подключение к сокету
func Chat(s *chatservice.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// ctx := r.Context()

		ws, err := s.Upgrader().Upgrade(w, r, nil)
		if err != nil {
			slog.Error(err.Error())
		}
		defer ws.Close()

		client := chatservice.NewClient(ws, s)
		s.Add(client)
		client.Listen()
	}
}
