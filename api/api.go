package api

import (
	"authorization-service/api/handler/oauth2"
	"authorization-service/api/handler/oidc"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/ory/fosite"
)

func NewMux(provider fosite.OAuth2Provider) *chi.Mux {
	var (
		mux = chi.NewMux()
	)

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)

	mux.Mount("/oauth2", oauth2.NewOauth2(provider))
	mux.Mount("/.well-known", oidc.NewOidc())

	return mux
}
