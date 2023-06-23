package store

import (
	"authorization-service/ent"
	"context"

	"github.com/ory/fosite"
	"golang.org/x/text/language"
)

func saveFositeRequest(ctx context.Context, client *ent.Client, req fosite.Requester, c *ent.Clients, s *ent.Session) (*ent.Request, error) {
	return client.Request.
		Create().
		SetID(req.GetID()).
		SetRequestedAt(req.GetRequestedAt()).
		SetClientID(c).
		SetScopes(req.GetRequestedScopes()).
		SetGrantedScopes(req.GetGrantedScopes()).
		SetForm(req.GetRequestForm()).
		SetRequestedAudience(req.GetRequestedAudience()).
		SetGrantedAudience(req.GetGrantedAudience()).
		SetSessionID(s).
		SetLang(language.AmericanEnglish).
		Save(ctx)
}
