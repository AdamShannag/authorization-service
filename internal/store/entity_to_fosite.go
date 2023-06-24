package store

import (
	"authorization-service/ent"

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
