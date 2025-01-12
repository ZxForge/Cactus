package core

import (
	"cactus/internal/error/validation"
	"cactus/internal/http/request"
	"cactus/internal/http/response"
	"context"
	"encoding/json"
	"net/http"
)

type abortEmailService interface {
	AbortMessage(ctx context.Context, uuid string) (string, error)
}

// Отмена рассылки писем по UUID что был передан для идентификации сообщения
func AbortMessage(s abortEmailService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req request.AbortMessageRequest
		_ = json.NewDecoder(r.Body).Decode(&req)
		errors, err := validation.ValidateStructure(&req)
		if err != nil {
			response.ResponseFailJSON(w, "Ошибка при проверке полей, проверьте структуру.")
			return
		}
		if errors != nil {
			response.ResponseValidationJSON(w, "Ошибка валидации, проверьте отправляемые поля", errors)
			return
		}

		// TODO обработать ошибку
		status, _ := s.AbortMessage(r.Context(), req.UUID)
		response.ResponseOKJSON(w, response.AbortMessageResponse{
			Status: status,
		})
	}
}
