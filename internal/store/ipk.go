package store

import (
	"authorization-service/ent"
	"context"

	"github.com/ory/fosite"
	"gopkg.in/square/go-jose.v2"
)

func getPublicKey(ctx context.Context, client *ent.Client, issuer string, subject string, keyId string) (*jose.JSONWebKey, error) {
	if pks, err := getOnePksFromIssuerSubjectByKey(client, ctx, issuer, subject, keyId); err != nil {
		return nil, err
	} else {
		return &pks.JSONWebKey, nil
	}
}

func getPublicKeys(ctx context.Context, client *ent.Client, issuer string, subject string) (*jose.JSONWebKeySet, error) {
	if pks, err := getAllPksFromIssuerBySubject(client, ctx, issuer, subject); err != nil || len(pks) == 0 {
		return nil, fosite.ErrNotFound
	} else {
		keys := make([]jose.JSONWebKey, 0, len(pks))
		for _, keyScopes := range pks {
			keys = append(keys, keyScopes.JSONWebKey)
		}
		return &jose.JSONWebKeySet{Keys: keys}, nil
	}
}

func getPublicKeyScopes(ctx context.Context, client *ent.Client, issuer string, subject string, keyId string) ([]string, error) {
	if pks, err := getOnePksFromIssuerSubjectByKey(client, ctx, issuer, subject, keyId); err != nil {
		return nil, err
	} else {
		return pks.Scopes, nil
	}
}
