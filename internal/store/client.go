package store

import (
	"authorization-service/ent"
	"authorization-service/ent/clients"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ory/fosite"
)

func getClient(ctx context.Context, client *ent.Client, id string) (fosite.Client, error) {
	c, err := client.Clients.
		Query().
		Where(clients.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, fosite.ErrNotFound
	}

	return toFositeClient(c), nil
}

func getClientFromRequest(client *ent.Client, req string, ctx context.Context) (*ent.Clients, error) {
	return client.Clients.Get(ctx, req)
}

func saveFositeSession(ctx context.Context, client *ent.Client, session fosite.Session, tokenType fosite.TokenType, sID ...string) (*ent.Session, error) {
	if len(sID) != 0 {
		se, err := client.Session.Get(ctx, sID[0])
		if err == nil {
			m := se.ExpiresAt
			m[string(tokenType)] = session.GetExpiresAt(fosite.TokenType(tokenType))
			return se.Update().
				SetUsername(session.GetUsername()).
				SetSubject(session.GetSubject()).
				SetExpiresAt(m).
				Save(ctx)
		}
	}

	newUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return client.Session.
		Create().
		SetID(newUUID.String()).
		SetUsername(session.GetUsername()).
		SetSubject(session.GetSubject()).
		SetExpiresAt(map[string]time.Time{string(tokenType): session.GetExpiresAt(fosite.TokenType(tokenType))}).
		Save(ctx)
}
