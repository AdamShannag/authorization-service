package store

import (
	"authorization-service/ent"
	"context"
	"fmt"
	"log"

	"github.com/ory/fosite"
)

func createPKCE(ctx context.Context, client *ent.Client, code string, req fosite.Requester) error {

	r, err := client.Request.Get(ctx, req.GetID())

	if err != nil {
		return err
	}

	u, err := client.PKCES.
		Create().
		SetID(code).
		SetRequestID(r).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating n pkce: %w", err)
	}
	log.Println("pkce was created: ", u)
	return nil
}

func findPKCEByCode(ctx context.Context, client *ent.Client, code string) (fosite.Requester, error) {

	pkce, err := client.PKCES.Get(ctx, code)

	if err != nil {
		return nil, fosite.ErrNotFound
	}

	r, err := pkce.QueryRequestID().WithClientID().WithSessionID().Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to find a request object: %w", err)
	}

	return &fosite.Request{
		ID:                r.ID,
		RequestedAt:       r.RequestedAt,
		Client:            toFositeClient(r.Edges.ClientID),
		Session:           toFositeSession(r.Edges.SessionID),
		RequestedScope:    r.Scopes,
		GrantedScope:      r.GrantedScopes,
		Form:              r.Form,
		RequestedAudience: r.RequestedAudience,
		GrantedAudience:   r.GrantedAudience,
		Lang:              r.Lang,
	}, nil
}

func deletePKCEByCode(ctx context.Context, client *ent.Client, code string) error {
	return client.PKCES.DeleteOneID(code).Exec(ctx)
}
