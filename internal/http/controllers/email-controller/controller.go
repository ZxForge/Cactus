package emailcontroller

import (
	validationerror "cactus/internal/error/validation-error"
	emailrequest "cactus/internal/http/request/email-request"
	emailresponse "cactus/internal/http/response/email-response"
	"cactus/internal/model"
	schemaApp "cactus/internal/schema"
	email_shema "cactus/internal/schema/email-shema"
	"log/slog"
	"time"

	"github.com/gorilla/schema"

	"context"
	"encoding/json"
	"net/http"
)

type emailService interface {
	GetStatus(ctx context.Context, uuid string) (string, error)
	Abort(ctx context.Context, uuid string) (string, error)
	CreateMessage(
		ctx context.Context,
		title, message string,
		subjects []string,
		priority *uint8,
		typeTemplate *string,
		data *map[string]string,
		theme *string,
		sendLater *time.Time,
		sameAttachment *bool,
		files []schemaApp.File,
	) (model.Message, error)
	SaveFiles(ctx context.Context, r *http.Request) ([]schemaApp.File, error)
	GetEmails(ctx context.Context, clientID int, processID int) ([]email_shema.Email, error)
}

// Отправка сообщения
func Send(s emailService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// var unmarshalErr *json.UnmarshalTypeError
		ctx := r.Context()

		err := r.ParseMultipartForm(32 << 20) // 32 МБ
		if err != nil {
			slog.Error(err.Error())
		}

		// TODO Вынести выше
		var decoder = schema.NewDecoder()

		err = r.ParseForm()
		if err != nil {
			// Handle error
		}

		var req emailrequest.SendRequest

		// r.PostForm is a map of our POST form values
		err = decoder.Decode(&req, r.PostForm)
		if err != nil {
			slog.Error("Ошибка декодирования: ", slog.String("message", err.Error()))
			return
		}

		//_ = json.NewDecoder(r.Body).Decode(&req)
		errors := validationerror.ValidateStructure(&req)
		if errors != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(403)
			jsonErrors, _ := json.Marshal(emailresponse.CreateResponseError("Ошибка валидации, проверьте отправляемые поля", errors))
			w.Write(jsonErrors)
			return
		}

		files, err := s.SaveFiles(ctx, r)
		if err != nil {
			slog.Error(err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			jsonErrors, _ := json.Marshal(emailresponse.CreateResponseError(err.Error(), nil))
			w.Write(jsonErrors)
			return
		}

		message, _ := s.CreateMessage(
			ctx,
			req.Title,
			req.Message,
			req.Subjects,
			req.Priority,
			req.Type,
			req.Data,
			req.Theme,
			req.SendLater,
			req.SameAttachment,
			files,
		)
		// TODO вызов сервис методов и преобразование в response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		jsonResp, _ := json.Marshal(emailresponse.CreateResponseOK(emailresponse.SendResponse{
			Status: "ok",
			UUID:   message.UUID,
		}))
		w.Write(jsonResp)
	}
}

// Получение статуса по UUID что был передан для идентификации рассылки
func GetStatus(s emailService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req emailrequest.GetStatusRequest
		_ = json.NewDecoder(r.Body).Decode(&req)
		errors := validationerror.ValidateStructure(&req)
		if errors != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(403)
			jsonErrors, _ := json.Marshal(emailresponse.CreateResponseError("Ошибка валидации, проверьте отправляемые поля", errors))
			w.Write(jsonErrors)
			return
		}

		// TODO обработать ошибку
		status, _ := s.GetStatus(r.Context(), req.UUID)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		jsonResp, _ := json.Marshal(emailresponse.CreateResponseOK(emailresponse.GetStatusResponse{
			Status: status,
		}))
		w.Write(jsonResp)
	}
}

// Отмена рассылки писем по UUID что был передан для идентификации сообщения
func Abort(s emailService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req emailrequest.AbortRequest
		_ = json.NewDecoder(r.Body).Decode(&req)
		errors := validationerror.ValidateStructure(&req)
		if errors != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(403)
			jsonErrors, _ := json.Marshal(emailresponse.CreateResponseError("Ошибка валидации, проверьте отправляемые поля", errors))
			w.Write(jsonErrors)
			return
		}

		// TODO обработать ошибку
		status, _ := s.Abort(r.Context(), req.UUID)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		jsonResp, _ := json.Marshal(emailresponse.CreateResponseOK(emailresponse.AbortResponse{
			Status: status,
		}))
		w.Write(jsonResp)
	}
}

// Генерирует разметку для просмотра как будет выглядеть сообщение, если email
func Render(s emailService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Получение ссылки на страницу просмотра рассылки в реальном времени
func GetLink(s emailService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Тест GetLink прошел успешно"))
	}
}

func GetEmails(s emailService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req emailrequest.GetEmailsRequest
		_ = json.NewDecoder(r.Body).Decode(&req)
		errors := validationerror.ValidateStructure(&req)
		if errors != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(403)
			jsonErrors, _ := json.Marshal(emailresponse.CreateResponseError("Ошибка валидации, проверьте отправляемые поля", errors))
			w.Write(jsonErrors)
			return
		}

		// TODO обработать ошибку
		emails, _ := s.GetEmails(r.Context(), req.ClientId, req.ProcessId)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		jsonResp, _ := json.Marshal(emailresponse.CreateResponseOK(emailresponse.GetEmailsResponse{
			Emails: emails,
		}))
		w.Write(jsonResp)
	}
}
