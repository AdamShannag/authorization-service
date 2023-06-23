package oauth2

import (
	"github.com/go-chi/chi/v5"
	"github.com/ory/fosite"
)

type Oauth2 struct {
	*chi.Mux
	oauth2 fosite.OAuth2Provider
}

func NewOauth2(provider fosite.OAuth2Provider) Oauth2 {
	h := Oauth2{
		Mux:    chi.NewMux(),
		oauth2: provider,
	}

	h.Get("/auth", h.Auth)
	h.Post("/auth", h.Auth)
	h.Post("/token", h.Token)
	h.Post("/revoke", h.Revoke)
	h.Post("/introspect", h.Introspect)

	return h
}
