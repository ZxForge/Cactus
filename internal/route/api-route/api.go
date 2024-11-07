package apiroute

import (
	emailcontroller "cactus/internal/http/controllers/email-controller"
	"cactus/internal/http/middleware"
	"cactus/internal/route/api-route/app"
	emailservice "cactus/internal/service/email-service"

	"github.com/go-chi/chi/v5"
)

func New(
	emailService *emailservice.Service,
) *chi.Mux {
	r := chi.NewRouter()

	appRoute := app.New(emailService)

	r.Mount("/app", appRoute)

	r.Group(func(r chi.Router) {
		r.Use(middleware.CheckDomainToken(emailService))

		r.Post("/email/send/", emailcontroller.Send(emailService))
		r.Post("/email/status/", emailcontroller.GetStatus(emailService))
		r.Post("/email/abort/", emailcontroller.Abort(emailService))
		r.Post("/email/render/", emailcontroller.Render(emailService))
		r.Post("/email/link/", emailcontroller.GetLink(emailService))
	})

	return r
}
