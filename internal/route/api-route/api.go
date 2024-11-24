package apiroute

import (
	emailcontroller "cactus/internal/http/controllers/email-controller"
	"cactus/internal/http/middleware"
	"cactus/internal/route/api-route/app"
	emailservice "cactus/internal/service/email-service"
	processservice "cactus/internal/service/process-service"

	"github.com/go-chi/chi/v5"
	cors "github.com/go-chi/cors"
)

func New(
	emailService *emailservice.Service,
	processService *processservice.Service,
) *chi.Mux {
	
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	  }))

	appRoute := app.New(emailService, processService)

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
