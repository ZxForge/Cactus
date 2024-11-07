package emailservice

import (
	filestorage "cactus/internal/file-storage"
	"cactus/internal/model"
	schemaApp "cactus/internal/schema"
	email_shema "cactus/internal/schema/email-shema"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type Storage interface {
	CreateMessage(
		ctx context.Context,
		domainId int64,
		uuid string,
		priority uint8,
		sendLater time.Time,
		value []byte, // TODO тут jsonb, а не string
	) (model.Message, error)
	GetStatus(ctx context.Context, uuid string) (string, error)
	Abort(ctx context.Context, uuid string) (string, error)
	GetClientByToken(ctx context.Context, token string) (model.Client, error)
	GetMessagesByClienIDAndProcessID(ctx context.Context, clientId int, processId int) ([]model.Message, error)
}

type Service struct {
	storage     Storage
	fileStorage *filestorage.FileStorage
}

func New(storage Storage, fileStorage *filestorage.FileStorage) *Service {
	return &Service{
		storage:     storage,
		fileStorage: fileStorage,
	}
}

func (s *Service) Abort(ctx context.Context, uuid string) (string, error) {
	// TODO: после добавления redis дополнить метод удалением из очереди в redis
	str, err := s.storage.Abort(ctx, uuid)
	return str, err
}

func (s *Service) GetStatus(ctx context.Context, uuid string) (string, error) {
	str, err := s.storage.GetStatus(ctx, uuid)
	return str, err
}

func (s *Service) CreateMessage(
	ctx context.Context,
	title, message string,
	subjects []string,
	priority *uint8,
	typeTemplate *string,
	data *map[string]string,
	theme *string,
	sendLater *time.Time,
	sameAttachment *bool,
	files []schemaApp.File, // TODO реализовать загрузку файлов
) (model.Message, error) {

	if priority == nil {
		var defaultPriority uint8 = 0
		priority = &defaultPriority
	}

	if sendLater == nil {
		defaultSendLater := time.Now()
		sendLater = &defaultSendLater
	}

	if theme == nil {
		defaultTheme := "default"
		theme = &defaultTheme
	}

	if data == nil {
		defaultData := make(map[string]string)
		data = &defaultData
	}

	if sameAttachment == nil {
		defaultSameAttachment := false
		sameAttachment = &defaultSameAttachment
	}

	var domainPriority = ctx.Value("domain-priority")

	if domainPriority != nil {
		complexPriority := *priority + domainPriority.(uint8)
		priority = &complexPriority
	}

	// TODO: после добавления redis дополнить метод удалением из очереди в redis

	value, _ := json.Marshal(email_shema.Email{
		Type:           *typeTemplate,
		Title:          title,
		Message:        message,
		Subjects:       subjects,
		Data:           *data,
		Theme:          *theme,
		SameAttachment: *sameAttachment,
	})

	newMessage, _ := s.storage.CreateMessage(
		ctx,
		1, // TODO реазиловать из контекста http
		uuid.New().String(),
		*priority,
		*sendLater,
		value,
	)
	return newMessage, nil
}

func (s *Service) GetClientByToken(ctx context.Context, token string) (model.Client, error) {
	domain, _ := s.storage.GetClientByToken(ctx, token) // TODO сделать обработку ошибки
	return domain, nil
}

func (s *Service) SaveFiles(ctx context.Context, r *http.Request) ([]schemaApp.File, error) {

	file, handler, err := r.FormFile("go")
	if err != nil {
		return []schemaApp.File{}, fmt.Errorf("%v", "Ошибка получения файла из тела запроса")
	}
	defer file.Close()

	info, err := s.fileStorage.Storage.PutObject(ctx, "cactus", handler.Filename, file, handler.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return []schemaApp.File{}, fmt.Errorf("%v", err.Error())
	}
	//dst, err := os.Create(handler.Filename)
	//if err != nil {
	//	return []schemaApp.File{}, fmt.Errorf("%v", "Ошибка создания файла")
	//}
	//defer dst.Close()

	//if _, err := io.Copy(dst, file); err != nil {
	//	slog.Error(err.Error())
	//	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(403)
	//	jsonErrors, _ := json.Marshal(emailresponse.CreateResponseError("500.errors.upload_image.cannot_copy_to_file", nil))
	//	w.Write(jsonErrors)
	//	return
	//}

	return []schemaApp.File{
		schemaApp.File{
			Title:      "картинка",
			PathToFile: info.Key,
		},
	}, nil
}

func (s *Service) GetEmails(ctx context.Context, clientID int, processID int) ([]email_shema.Email, error) {
	messages, err := s.storage.GetMessagesByClienIDAndProcessID(ctx, clientID, processID)
	if err != nil {
		return []email_shema.Email{}, err
	}

	var emails []email_shema.Email = []email_shema.Email{}

	for _, message := range messages {
		var email email_shema.Email
		email.Title = message.Title
		emails = append(emails, email)
	}

	return emails, nil
}
