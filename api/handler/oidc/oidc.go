package oidc

import (
	"github.com/go-chi/chi/v5"
)

type Oidc struct {
	*chi.Mux
	// add extra members here, such as your repo or client
}

func NewOidc(extras ...any) Oidc {
	h := Oidc{
		Mux: chi.NewMux(),
		// initialize extra members here, such as your repo or client
	}

	h.Get("/openid-configuration", h.OpenIDConfigurationEndpoint)
	h.Get("/jwks.json", h.JSONWebKeysEndpoint)
	return h
}
