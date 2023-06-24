package store

import (
	"authorization-service/ent"
	"context"
	"log"
	"time"

	"github.com/ory/fosite"
	"golang.org/x/text/language"
)

func createIDSession(ctx context.Context, client *ent.Client, authorizeCode string, req fosite.Requester) error {
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

	u, err := client.IDSessions.
		Create().
		SetID(authorizeCode).
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

	log.Printf("\n\n==============================\nREQUEST CreateOpenIDConnectSession: \n%+v\nCODE: %s\n==============================", req, authorizeCode)
	log.Printf("\n\n==============================\nCreated IDSESSION:%+v\n\n==============================", u)
	log.Printf("\n\n==============================\nCreated SESSION IDSESSION:%+v\n\n==============================", s)

	return nil
}

func findIDSessionByCode(ctx context.Context, client *ent.Client, code string, req fosite.Requester) (fosite.Requester, error) {

	tokenTypes := []fosite.TokenType{fosite.AuthorizeCode, fosite.AccessToken, fosite.RefreshToken, fosite.IDToken, fosite.PushedAuthorizeRequestContext}

	m := map[string]time.Time{}

	for _, s := range tokenTypes {
		if req.GetSession().GetExpiresAt(s).After(time.Now()) {
			m[string(s)] = req.GetSession().GetExpiresAt(s)
		}
	}

	iDSession, err := client.IDSessions.Get(ctx, code)

	if err != nil {
		return nil, fosite.ErrNotFound
	}

	c, err := iDSession.QueryClientID().Only(ctx)

	if err != nil {
		return nil, fosite.ErrNotFound
	}

	sessionEnt, err := client.Session.Get(ctx, req.GetID())
	session := &fosite.DefaultSession{}
	if err == nil {
		session = &fosite.DefaultSession{
			ExpiresAt: getSessionExpiryAtFositeMap(sessionEnt.ExpiresAt),
			Username:  sessionEnt.Username,
			Subject:   sessionEnt.Subject,
			Extra:     sessionEnt.Extra,
		}

		iDSession, err = iDSession.Update().SetSessionID(sessionEnt).SetRequestID(sessionEnt.ID).Save(ctx)
		if err != nil {
			return nil, fosite.ErrServerError
		}

	} else {
		sessionEnt, err := client.Session.Get(ctx, iDSession.RequestID)

		if err != nil {
			return nil, fosite.ErrNotFound
		}
		session = &fosite.DefaultSession{
			ExpiresAt: getSessionExpiryAtFositeMap(sessionEnt.ExpiresAt),
			Username:  sessionEnt.Username,
			Subject:   sessionEnt.Subject,
			Extra:     sessionEnt.Extra,
		}
	}

	fSe, ok := sessionEnt.Session.(*fosite.Session)

	if ok {
		log.Println(fSe)
	}

	idSession := &fosite.Request{
		ID:                iDSession.ID,
		RequestedAt:       iDSession.RequestedAt,
		Client:            toFositeClient(c),
		Session:           session,
		RequestedScope:    iDSession.Scopes,
		GrantedScope:      iDSession.GrantedScopes,
		Form:              iDSession.Form,
		RequestedAudience: iDSession.RequestedAudience,
		GrantedAudience:   iDSession.GrantedAudience,
		Lang:              iDSession.Lang,
	}

	log.Printf("\n\nREQUEST GetOpenIDConnectSession: \n%+v\nCODE: %s\n\n==============================", req, code)
	log.Printf("\n\nFOUND IDSESSION: %+v\n==============================", idSession)
	log.Printf("\n\nFOUND SESSION IDSESSION: %+v\n==============================", idSession.GetSession())

	return idSession, nil
}

func deleteIDSessionByCode(ctx context.Context, client *ent.Client, code string) error {
	return client.IDSessions.DeleteOneID(code).Exec(ctx)
}
