package store

import (
	"authorization-service/ent"
	"time"

	"github.com/ory/fosite"
)

func toFositeClient(c *ent.Clients) *fosite.DefaultClient {
	return &fosite.DefaultClient{
		ID:             c.ID,
		Secret:         c.ClientSecret,
		RotatedSecrets: c.RotatedSecrets,
		RedirectURIs:   c.RedirectUris,
		GrantTypes:     c.GrantTypes,
		ResponseTypes:  c.ResponseTypes,
		Scopes:         c.Scopes,
		Public:         c.Public,
	}
}

func toFositeSession(s *ent.Session) *fosite.DefaultSession {

	expiredAt := make(map[fosite.TokenType]time.Time)
	for k, v := range s.ExpiresAt {
		expiredAt[fosite.TokenType(k)] = v
	}

	return &fosite.DefaultSession{
		ExpiresAt: expiredAt,
		Username:  s.Username,
		Subject:   s.Subject,
		Extra:     s.Extra,
	}
}
