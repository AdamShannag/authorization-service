package store

import (
	"authorization-service/ent"
	"context"
	"fmt"
	"log"

	"github.com/ory/fosite"
)

func createIDSession(ctx context.Context, client *ent.Client, authorizeCode string, req fosite.Requester) error {

	r, err := client.Request.Get(ctx, req.GetID())

	if err != nil {
		return err
	}

	// m := r.Form
	// for k, v := range req.GetRequestForm() {
	// 	m[k] = v
	// }

	err = r.Update().SetForm(req.GetRequestForm()).Exec(ctx)

	if err != nil {
		return err
	}

	u, err := client.IDSessions.
		Create().
		SetID(authorizeCode).
		SetRequestID(r).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating an id_session: %w", err)
	}
	log.Println("id_session was created: ", u)
	return nil
}

func findIDSessionByCode(ctx context.Context, client *ent.Client, code string) (fosite.Requester, error) {

	iDSession, err := client.IDSessions.Get(ctx, code)

	if err != nil {
		return nil, fosite.ErrNotFound
	}

	r, err := iDSession.QueryRequestID().WithSessionID().WithClientID().Only(ctx)

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

func deleteIDSessionByCode(ctx context.Context, client *ent.Client, code string) error {
	return client.IDSessions.DeleteOneID(code).Exec(ctx)
}
