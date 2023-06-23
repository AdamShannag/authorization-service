package store

import (
	"authorization-service/ent"
	"context"
	"fmt"
	"log"

	"github.com/ory/fosite"
)

func createRefreshToken(ctx context.Context, client *ent.Client, signature string, req fosite.Requester) error {
	r, err := client.Request.Get(ctx, req.GetID())

	if err != nil {
		return err
	}

	se, err := r.QuerySessionID().Only(ctx)

	if err != nil {
		log.Println(err)
		return err
	}

	_, err = saveFositeSession(ctx, client, req.GetSession(), fosite.RefreshToken, se.ID)

	if err != nil {
		log.Println(err)
		return err
	}

	u, err := client.RefreshTokens.
		Create().
		SetID(signature).
		SetActive(true).
		SetRequestID(r).
		Save(ctx)

	if err != nil {
		return fmt.Errorf("failed creating an authorize_code: %w", err)
	}
	log.Println("refresh_token was created: ", u)
	return nil
}

func findRefreshTokenByID(ctx context.Context, client *ent.Client, signature string) (fosite.Requester, bool, error) {

	token, err := client.RefreshTokens.Get(ctx, signature)

	if err != nil {
		return nil, token.Active, fmt.Errorf("failed to find an refresh_token: %w", err)
	}

	r, err := token.QueryRequestID().WithClientID().WithSessionID().Only(ctx)

	if err != nil {
		return nil, token.Active, fmt.Errorf("failed to find a request object: %w", err)
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
	}, token.Active, nil
}

func deleteRefreshTokenByID(ctx context.Context, client *ent.Client, signature string) error {
	return client.RefreshTokens.DeleteOneID(signature).Exec(ctx)
}

func updateRefreshTokenStatusByID(ctx context.Context, client *ent.Client, signature string, status bool) error {

	err := client.RefreshTokens.UpdateOneID(signature).SetActive(status).Exec(ctx)

	if err != nil {
		return fosite.ErrNotFound
	}

	return nil
}
