package store

import (
	"authorization-service/ent"
	"authorization-service/ent/user"
	"context"

	"github.com/ory/fosite"
)

func authenticateUser(ctx context.Context, client *ent.Client, name, secret string) error {
	user, err := client.User.Query().Where(user.Username(name)).Only(ctx)

	if err != nil {
		return fosite.ErrNotFound
	}

	if user.Password != secret {
		return fosite.ErrNotFound.WithDebug("Invalid credentials")
	}

	return nil
}
