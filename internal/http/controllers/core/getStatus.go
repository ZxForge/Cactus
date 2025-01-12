package core

import (
	"cactus/internal/error/validation"
	"cactus/internal/http/request"
	"cactus/internal/http/response"
	"context"
	"encoding/json"
	"net/http"
)

type getStatusService interface {
	GetStatus(ctx context.Context, uuid string) (string, error)
}

// Получение статуса по UUID что был передан для идентификации рассылки
func GetStatus(s getStatusService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req request.GetStatusMessageRequest
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

		status, err := s.GetStatus(r.Context(), req.UUID)
		if err != nil {
			// TODO: добавить trace для ошибок
			response.ResponseFailJSON(w, "Внутренняя ошибка сервера. Пожалуйста, попробуйте позже.")
		}

		response.ResponseOKJSON(w, response.GetStatusMessageResponse{
			Status: status,
		})
	}
}
