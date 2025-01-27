package main

import (
	"cactus/internal/logger"
	configschema "cactus/internal/pkg/configSchema"
	"fmt"
	"log/slog"
	"net/smtp"

	rdb "cactus/internal/storage/redis"

	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

// type EmailData struct {
// 	To      string `json:"to"`
// 	Subject string `json:"subject"`
// 	Body    string `json:"body"`
// }

type SMTPWorkerConfig struct {
	Host string `validate:"required,hostname" slug:"host"`
	Port string `validate:"required,numeric" slug:"port"`
	From string `validate:"required,email" slug:"from"`

	smtp *smtp.Client
}

func (conf *SMTPWorkerConfig) Update(values map[string]interface{}) error {
	backup := *conf

	MapToStruct(values, conf)

	validate := validator.New()
	if err := validate.Struct(conf); err != nil {
		*conf = backup
	}

	conf.smtp = &smtp.Client{}

	return nil
}

func (conf *SMTPWorkerConfig) Send() {
	smtpServer := "localhost:1025"
	from := "sender@example.com"            // Отправитель
	to := []string{"recipient@example.com"} // Получатель

	// Формирование сообщения
	subject := "Subject: Тестовое письмо\n"
	body := "Это тестовое письмо, отправленное через MailHog.\n"
	messageb := []byte(subject + "\n" + body)

	// Отправка письма
	err := smtp.SendMail(smtpServer, nil, from, to, messageb)
	if err != nil {
		fmt.Println("Ошибка отправки письма:", err)
		return
	}
}

func (conf *SMTPWorkerConfig) String() string {
	return fmt.Sprintf("address:%v:%v form:%v", conf.Host, conf.Port, conf.From)
}

func main() {
	Type := "email"
	Kind := "smtp"

	conf := MustLoad()

	if conf.WorkerUUID == "" {
		slog.Error("worker обязан иметь ID (UUID)")
		return
	}

	SMTPWorker := SMTPWorkerConfig{}

	ctx := context.Background()

	slog.SetDefault(logger.SetupLogger(conf.Env))

	RDBStorage, err := rdb.New(ctx, &redis.Options{
		Addr:     conf.Redis.Address,
		Password: conf.Redis.Password,
		Username: conf.Redis.User,
	})

	if err != nil {
		slog.Error("Ошибка создания соединения с redis: ", slog.Any("error", err.Error()))
		return
	}

	defer func() {
		err := RDBStorage.Close()
		if err != nil {
			slog.Error("Ошибка закрытия соединения с redis:", slog.Any("error", err.Error()))
		}
	}()

	worker := NewWorker(ctx, RDBStorage, WorkerConfig{
		Token:      conf.Token,
		WorkerKind: Kind,
		WorkerType: Type,
		WorkerUUID: conf.WorkerUUID,
		ConfigSchema: []configschema.ConfigField{
			{
				Type: "host",
				Slug: "host",
				Name: "ip адрес сервера SMTP",
			},
			{
				Type: "numeric",
				Slug: "port",
				Name: "Порт сервера",
			},
			{
				Type: "email",
				Slug: "from",
				Name: "Адрем отправителя",
			},
		},
	})

	worker.SetConfigHandler(func(message Message) {
		SMTPWorker.Update(message.Value)
		fmt.Printf("Configuring worker with task: %+v\n", message)
	})

	worker.SetHandler(func(message QueueMessage) {
		fmt.Println(message)
		fmt.Println(SMTPWorker.String())
	})

	fmt.Println("Воркер запущен")

	// TODO добавить CTRL+C сигнал и graceful-shutdown
	worker.Run()

	fmt.Println("Воркер остановлен")
}
