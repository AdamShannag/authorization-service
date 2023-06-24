package store

import (
	"authorization-service/ent"
	"context"
	"log"
	"time"

	"github.com/ory/fosite"
	"golang.org/x/text/language"
)

func createAccessToken(ctx context.Context, client *ent.Client, signature string, req fosite.Requester) error {
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

	_, err = client.AccessTokens.
		Create().
		SetID(signature).
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

func findAccessTokenBySignature(ctx context.Context, client *ent.Client, signature string) (fosite.Requester, error) {

	token, err := client.AccessTokens.Get(ctx, signature)

	if err != nil {
		return nil, fosite.ErrNotFound
	}

	c, err := token.QueryClientID().Only(ctx)

	if err != nil {
		return nil, fosite.ErrNotFound
	}

	sessionEnt, err := client.Session.Get(ctx, token.RequestID)

	if err != nil {
		return nil, fosite.ErrNotFound
	}

	fSe,ok := sessionEnt.Session.(*fosite.Session)

	if ok{
		log.Println(fSe)
	}


	session := &fosite.DefaultSession{
		ExpiresAt: getSessionExpiryAtFositeMap(sessionEnt.ExpiresAt),
		Username:  sessionEnt.Username,
		Subject:   sessionEnt.Subject,
		Extra:     sessionEnt.Extra,
	}

	return &fosite.Request{
		ID:                token.ID,
		RequestedAt:       token.RequestedAt,
		Client:            toFositeClient(c),
		Session:           session,
		RequestedScope:    token.Scopes,
		GrantedScope:      token.GrantedScopes,
		Form:              token.Form,
		RequestedAudience: token.RequestedAudience,
		GrantedAudience:   token.GrantedAudience,
		Lang:              token.Lang,
	}, nil
}

func deleteAccessTokenBySignature(ctx context.Context, client *ent.Client, signature string) error {
	return client.AccessTokens.DeleteOneID(signature).Exec(ctx)
}
