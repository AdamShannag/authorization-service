package store

import (
	"authorization-service/ent"
	"context"
	"log"
	"time"

	"github.com/ory/fosite"
	"golang.org/x/text/language"
)

func createPKCE(ctx context.Context, client *ent.Client, code string, req fosite.Requester) error {
	tokenTypes := []fosite.TokenType{fosite.AuthorizeCode, fosite.AccessToken, fosite.RefreshToken, fosite.IDToken, fosite.PushedAuthorizeRequestContext}

	m := map[string]time.Time{}

	for _, s := range tokenTypes {
		if req.GetSession().GetExpiresAt(s).After(time.Now()) {
			m[string(s)] = req.GetSession().GetExpiresAt(s)
		}
	}

	s, err := client.Session.Get(ctx, req.GetID())
	if err == nil {
		s, err = client.Session.UpdateOneID(req.GetID()).Save(ctx)

		if err != nil {
			return fosite.ErrServerError
		}
	} else {
		s, err = client.Session.Create().
			SetID(req.GetID()).
			SetExpiresAt(m).
			SetUsername(req.GetSession().GetUsername()).
			SetSubject(req.GetSession().GetSubject()).
			SetSession(req.GetSession()).
			Save(ctx)

		if err != nil {
			return fosite.ErrServerError
		}
	}

	_, err = client.PKCES.
		Create().
		SetID(code).
		SetRequestID(req.GetID()).
		SetRequestedAt(req.GetRequestedAt()).
		SetClientIDID(req.GetClient().GetID()).
		SetSessionID(s).
		SetScopes(req.GetRequestedScopes()).
		SetGrantedScopes(req.GetGrantedScopes()).
		SetForm(req.GetRequestForm()).
		SetRequestedAudience(req.GetRequestedAudience()).
		SetGrantedAudience(req.GetGrantedAudience()).
		SetLang(language.English).
		Save(ctx)
	if err != nil {
		return fosite.ErrServerError
	}
	return nil
}

func findPKCEByCode(ctx context.Context, client *ent.Client, code string) (fosite.Requester, error) {

	pkce, err := client.PKCES.Get(ctx, code)

	if err != nil {
		return nil, fosite.ErrNotFound
	}

	c, err := pkce.QueryClientID().Only(ctx)

	if err != nil {
		return nil, fosite.ErrNotFound
	}

	sessionEnt, err := client.Session.Get(ctx, pkce.RequestID)

	if err != nil {
		return nil, fosite.ErrNotFound
	}

	session := &fosite.DefaultSession{
		ExpiresAt: getSessionExpiryAtFositeMap(sessionEnt.ExpiresAt),
		Username:  sessionEnt.Username,
		Subject:   sessionEnt.Subject,
		Extra:     sessionEnt.Extra,
	}

	fSe, ok := sessionEnt.Session.(*fosite.Session)

	if ok {
		log.Println(fSe)
	}

	return &fosite.Request{
		ID:                pkce.ID,
		RequestedAt:       pkce.RequestedAt,
		Client:            toFositeClient(c),
		Session:           session,
		RequestedScope:    pkce.Scopes,
		GrantedScope:      pkce.GrantedScopes,
		Form:              pkce.Form,
		RequestedAudience: pkce.RequestedAudience,
		GrantedAudience:   pkce.GrantedAudience,
		Lang:              pkce.Lang,
	}, nil
}

func deletePKCEByCode(ctx context.Context, client *ent.Client, code string) error {
	return client.PKCES.DeleteOneID(code).Exec(ctx)
}
