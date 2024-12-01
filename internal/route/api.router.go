package route

import (
	email_controller "cactus/internal/http/controllers/email"
	"cactus/internal/http/middleware"
	email_service "cactus/internal/service/email"

	"github.com/go-chi/chi/v5"
)

func addRouteApi(
	r *chi.Mux,
	emailService *email_service.Service,
) *chi.Mux {
	r.Route("/app", func(r chi.Router) {
		r.Use(middleware.AuthSession())

		r.Route("/email", func(r chi.Router) {
			r.Post("/list/", email_controller.GetEmails(emailService))
		})
	})

	r.Group(func(r chi.Router) {
		r.Use(middleware.CheckDomainToken(emailService))

		r.Post("/email/send/", email_controller.Send(emailService))
		r.Post("/email/status/", email_controller.GetStatus(emailService))
		r.Post("/email/abort/", email_controller.Abort(emailService))
		r.Post("/email/render/", email_controller.Render(emailService))
		r.Post("/email/link/", email_controller.GetLink(emailService))
	})

	return r
}
