package store

import (
	"authorization-service/ent"

	"github.com/ory/fosite"
)

func toEntityRequest(r fosite.Requester) *ent.Request {

	e := &ent.Request{
		ID:                r.GetID(),
		RequestedAt:       r.GetRequestedAt(),
		Scopes:            r.GetRequestedScopes(),
		GrantedScopes:     r.GetGrantedScopes(),
		RequestedAudience: r.GetRequestedAudience(),
		GrantedAudience:   r.GetGrantedAudience(),
		Form:              r.GetRequestForm(),
	}

	return e
}
