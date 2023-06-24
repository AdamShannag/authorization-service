package store

import (
	"authorization-service/ent"
	"context"
	"log"
	"time"

	"github.com/ory/fosite"
	"golang.org/x/text/language"
)

func createAuthorizeCode(ctx context.Context, client *ent.Client, code string, req fosite.Requester) error {
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

	_, err = client.AuthorizeCodes.
		Create().
		SetID(code).
		SetActive(true).
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

func getSessionExpiryAtEntMap(tokenType fosite.TokenType, session fosite.Session) map[string]time.Time {
	return map[string]time.Time{string(tokenType): session.GetExpiresAt(fosite.TokenType(tokenType))}
}

func getSessionExpiryAtFositeMap(entMap map[string]time.Time) map[fosite.TokenType]time.Time {
	fositeMap := make(map[fosite.TokenType]time.Time)
	for k, v := range entMap {
		fositeMap[fosite.TokenType(k)] = v
	}
	return fositeMap
}

func findAuthorizeCodeByID(ctx context.Context, client *ent.Client, code string) (fosite.Requester, bool, error) {

	authCode, err := client.AuthorizeCodes.Get(ctx, code)
	if err != nil {
		return nil, authCode.Active, fosite.ErrNotFound
	}

	c, err := authCode.QueryClientID().Only(ctx)

	if err != nil {
		return nil, authCode.Active, fosite.ErrNotFound
	}

	sessionEnt, err := client.Session.Get(ctx, authCode.RequestID)

	if err != nil {
		return nil, authCode.Active, fosite.ErrNotFound
	}

	fSe, ok := sessionEnt.Session.(*fosite.Session)

	if ok {
		log.Println(fSe)
	}

	session := &fosite.DefaultSession{
		ExpiresAt: getSessionExpiryAtFositeMap(sessionEnt.ExpiresAt),
		Username:  sessionEnt.Username,
		Subject:   sessionEnt.Subject,
		Extra:     sessionEnt.Extra,
	}

	return &fosite.Request{
		ID:                authCode.ID,
		RequestedAt:       authCode.RequestedAt,
		Client:            toFositeClient(c),
		Session:           session,
		RequestedScope:    authCode.Scopes,
		GrantedScope:      authCode.GrantedScopes,
		Form:              authCode.Form,
		RequestedAudience: authCode.RequestedAudience,
		GrantedAudience:   authCode.GrantedAudience,
		Lang:              authCode.Lang,
	}, authCode.Active, nil
}

func updateAuthorizeCodeStatusByID(ctx context.Context, client *ent.Client, code string, status bool) error {

	_, err := client.AuthorizeCodes.UpdateOneID(code).SetActive(status).Save(ctx)

	if err != nil {
		return fosite.ErrNotFound
	}

	return nil
}
