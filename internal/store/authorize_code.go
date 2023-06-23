package store

import (
	"authorization-service/ent"
	"context"
	"fmt"
	"log"

	"github.com/ory/fosite"
)

func createAuthorizeCode(ctx context.Context, client *ent.Client, code string, req fosite.Requester) error {

	c, err := getClientFromRequest(client, req.GetClient().GetID(), ctx)

	if err != nil {
		// log.Println(err)
		return err
	}

	s, err := saveFositeSession(ctx, client, req.GetSession(), fosite.AuthorizeCode)

	if err != nil {
		// log.Println(err)

		return err
	}

	r, err := saveFositeRequest(ctx, client, req, c, s)
	log.Println("\n\nREQ: ", req)
	log.Println("\n\nR: ", r)
	if err != nil {
		// log.Println(err)

		return err
	}

	u, err := client.AuthorizeCodes.
		Create().
		SetID(code).
		SetActive(true).
		SetRequestID(r).
		Save(ctx)
	if err != nil {
		// log.Println(err)
		return fmt.Errorf("failed creating an authorize_code: %w", err)
	}
	log.Println("authorize_code was created: ", u)
	return nil
}

func findAuthorizeCodeByID(ctx context.Context, client *ent.Client, code string) (fosite.Requester, bool, error) {

	authCode, err := client.AuthorizeCodes.Get(ctx, code)
	if err != nil {
		// log.Println(err, "failed to find an authorize_code")
		return nil, authCode.Active, fmt.Errorf("failed to find an authorize_code: %w", err)
	}

	r, err := authCode.QueryRequestID().WithClientID().WithSessionID().Only(ctx)

	if err != nil {
		return nil, authCode.Active, fmt.Errorf("failed to find a request object: %w", err)
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
	}, authCode.Active, nil
}

func updateAuthorizeCodeStatusByID(ctx context.Context, client *ent.Client, code string, status bool) error {

	_, err := client.AuthorizeCodes.UpdateOneID(code).SetActive(status).Save(ctx)

	if err != nil {
		return fosite.ErrNotFound
	}

	return nil
}
