package app

import (
	emailcontroller "cactus/internal/http/controllers/email-controller"
	"cactus/internal/http/middleware"
	emailservice "cactus/internal/service/email-service"

	"github.com/go-chi/chi/v5"
)

func New(
	emailService *emailservice.Service,
) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.AuthSession())

	r.Route("/email", func(r chi.Router) {
		r.Post("/list/", emailcontroller.GetEmails(emailService))
	})

	return r
}
