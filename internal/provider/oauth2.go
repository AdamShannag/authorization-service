package provider

import (
	"authorization-service/ent"
	"authorization-service/internal/store"
	"crypto/rand"
	"crypto/rsa"
	"time"

	"github.com/ory/fosite"

	"github.com/ory/fosite/compose"
	"github.com/ory/fosite/handler/openid"
	"github.com/ory/fosite/token/jwt"
)

func NewOauth2Provider(client *ent.Client) fosite.OAuth2Provider {
	var (
		secret        = []byte("some-cool-secret-that-is-32bytes")
		store         = store.NewDBStore(client)
		privateKey, _ = rsa.GenerateKey(rand.Reader, 2048)

		config = &fosite.Config{
			AccessTokenLifespan: time.Minute * 30,
			GlobalSecret:        secret,
		}
	)

	oauth2 := compose.ComposeAllEnabled(config, store, privateKey)
	return oauth2
}

func NewSession(user string) *openid.DefaultSession {
	return &openid.DefaultSession{
		Claims: &jwt.IDTokenClaims{
			Issuer:      "https://fosite.my-application.com",
			Subject:     user,
			Audience:    []string{"https://my-client.my-application.com"},
			ExpiresAt:   time.Now().Add(time.Hour * 6),
			IssuedAt:    time.Now(),
			RequestedAt: time.Now(),
			AuthTime:    time.Now(),
		},
		Headers: &jwt.Headers{
			Extra: make(map[string]interface{}),
		},
	}
}
