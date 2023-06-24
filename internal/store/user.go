package store

import (
	"authorization-service/ent"
	"context"

	"github.com/ory/fosite"
)

func authenticateUser(ctx context.Context, client *ent.Client, name, secret string) error {
	user, err := client.User.Get(ctx, name)

	if err != nil {
		return fosite.ErrNotFound
	}

	if user.Password != secret {
		return fosite.ErrNotFound.WithDebug("Invalid credentials")
	}

	return nil
}
