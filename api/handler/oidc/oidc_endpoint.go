package oidc

import (
	"encoding/json"
	"net/http"
)

func (p *Oidc) OpenIDConfigurationEndpoint(w http.ResponseWriter, r *http.Request) {
	config := OpenIDConfiguration{
		Issuer:                             "http://localhost:8000",
		AuthorizationEndpoint:              "http://localhost:8000/oauth2/auth",
		TokenEndpoint:                      "http://localhost:8000/oauth2/token",
		UserInfoEndpoint:                   "http://localhost:8000/userinfo",
		JwksURI:                            "http://localhost:8000/.well-known/jwks.json",
		ResponseTypesSupported:             []string{"code", "token", "id_token"},
		SubjectTypesSupported:              []string{"public"},
		IDTokenSigningAlgValuesSupported:   []string{"RS256"},
		ScopesSupported:                    []string{"openid", "profile", "email"},
		ClaimsSupported:                    []string{"sub", "iss", "name", "email"},
		GrantTypesSupported:                []string{"authorization_code", "refresh_token"},
		TokenEndpointAuthMethodsSupported:  []string{"client_secret_basic"},
		RevocationEndpoint:                 "http://localhost:8000/oauth2/revoke",
		EndSessionEndpoint:                 "http://localhost:8000/oauth2/logout",
		IntrospectionEndpoint:              "http://localhost:8000/oauth2/introspect",
		CheckSessionIframe:                 "http://localhost:8000/oauth2/session/check",
		FrontChannelLogoutSupported:        true,
		FrontChannelLogoutSessionSupported: true,
		BackChannelLogoutSupported:         true,
		BackChannelLogoutSessionSupported:  true,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config)
}
