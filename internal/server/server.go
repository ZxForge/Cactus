package server

import (
	"cactus/internal/config"
	"cactus/internal/db"
	dbstorage "cactus/internal/db-storage"
	filestorage "cactus/internal/file-storage"
	apiroute "cactus/internal/route/api-route"
	wsroute "cactus/internal/route/ws-route"
	chatservice "cactus/internal/service/chat-service"
	emailservice "cactus/internal/service/email-service"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	db  *sqlx.DB
	app *http.Server
}

type CreateServerOption struct {
	Db     config.Database
	Server config.HTTPServer
	IsDev  bool
}

func Create(options CreateServerOption) (Server, error) {

	database, err := db.New(
		options.Db.Host,
		options.Db.Port,
		options.Db.Name,
		options.Db.User,
		options.Db.Pass,
	)
	if err != nil {
		return Server{}, fmt.Errorf("create database: %w", err)
	}

	// TODO вынести в конфиг
	endpoint := "local.work.ru:9000"

	accessKeyID := "H2opu9nTT9JCVaiFer0o"
	secretAccessKey := "MhO1IZ7vOQB5lzaU6Rg4nPBO8NnrFSDUo5eXi47Z"
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		slog.Error(err.Error())
	}

	dbStorage := dbstorage.New(database)
	fileStorage := filestorage.New(minioClient)

	// appService := appservice.New(dbStorage)
	emailService := emailservice.New(dbStorage, fileStorage)
	chatServer := chatservice.NewService()
	// TODO: сделать общий обработчик ошибок на уровне middleware для http(s)
	r := chi.NewRouter()
	apiRoute := apiroute.New(
		emailService,
	)
	wsRoute := wsroute.New(
		chatServer,
	)

	r.Mount("/api", apiRoute)
	r.Mount("/ws", wsRoute)

	app := &http.Server{
		Addr:         options.Server.Address,
		Handler:      r,
		IdleTimeout:  options.Server.IdleTimeout,
		ReadTimeout:  options.Server.Timeout,
		WriteTimeout: options.Server.Timeout,
	}

	return Server{
		db:  database,
		app: app,
	}, nil
}

func (s *Server) Start() error {
	defer s.db.Close()

	dieServer := make(chan error, 1)
	defer close(dieServer)

	go func() {
		if err := s.app.ListenAndServe(); err != nil {
			dieServer <- err
		}
	}()

	err := <-dieServer
	return err
}
