package store

import (
	"authorization-service/ent"
	"context"
	"fmt"
	"log"

	"github.com/ory/fosite"
)

func createAccessToken(ctx context.Context, client *ent.Client, signature string, req fosite.Requester) error {

	r, err := client.Request.Get(ctx, req.GetID())

	log.Println("\n\nREQ: ", req)
	log.Println("\n\nR: ", r)

	if err != nil {
		return err
	}

	se, err := r.QuerySessionID().Only(ctx)

	if err != nil {
		// log.Println(err)
		return err
	}

	_, err = saveFositeSession(ctx, client, req.GetSession(), fosite.AccessToken, se.ID)

	if err != nil {
		// log.Println(err)
		return err
	}

	u, err := client.AccessTokens.
		Create().
		SetID(signature).
		SetRequestID(r).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating an access_token: %w", err)
	}
	log.Println("access_token was created: ", u)
	return nil
}

func findAccessTokenBySignature(ctx context.Context, client *ent.Client, signature string) (fosite.Requester, error) {

	token, err := client.AccessTokens.Get(ctx, signature)

	if err != nil {
		return nil, fosite.ErrNotFound
	}

	r, err := token.QueryRequestID().WithSessionID().WithClientID().Only(ctx)

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

func deleteAccessTokenBySignature(ctx context.Context, client *ent.Client, signature string) error {
	return client.AccessTokens.DeleteOneID(signature).Exec(ctx)
}
