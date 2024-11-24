package app

import (
	emailcontroller "cactus/internal/http/controllers/email-controller"
	processController "cactus/internal/http/controllers/processes"
	"cactus/internal/http/middleware"
	emailservice "cactus/internal/service/email-service"
	processservice "cactus/internal/service/process-service"

	"github.com/go-chi/chi/v5"
)

func New(
	emailService *emailservice.Service,
	processService *processservice.Service,
) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.AuthSession())
	r.Use()

	r.Route("/email", func(r chi.Router) {
		r.Post("/list/", emailcontroller.GetEmails(emailService))
	})

	r.Route("/process", func(r chi.Router) {
		r.Post("/list/", processController.GetProcessesFor(processService))
	})

	return r
}
