package store

import (
	"authorization-service/ent"
	"context"

	"github.com/ory/fosite"
)

func getClient(ctx context.Context, client *ent.Client, id string) (fosite.Client, error) {
	c, err := client.Clients.Get(ctx, id)

	if err != nil {
		return nil, fosite.ErrNotFound
	}

	return toFositeClient(c), nil
}
