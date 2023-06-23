package store

import (
	"authorization-service/ent"
	"authorization-service/ent/publickeyscopes"
	"authorization-service/ent/subjectpublickeys"
	"context"

	"github.com/ory/fosite"
)

func getOnePksFromIssuerSubjectByKey(client *ent.Client, ctx context.Context, issuer string, subject string, keyId string) (*ent.PublicKeyScopes, error) {
	ipk, err := getIPK(client, ctx, issuer)
	if err != nil {
		return nil, err
	}

	spk, err := getSpkFromIpk(ipk, subject, ctx)
	if err != nil {
		return nil, err
	}

	pks, err := getPksFromSpk(spk, keyId, ctx)
	if err != nil {
		return nil, err
	}

	return pks, nil
}
func getAllPksFromIssuerBySubject(client *ent.Client, ctx context.Context, issuer string, subject string) ([]*ent.PublicKeyScopes, error) {
	ipk, err := getIPK(client, ctx, issuer)
	if err != nil {
		return nil, err
	}

	spk, err := getSpkFromIpk(ipk, subject, ctx)
	if err != nil {
		return nil, err
	}

	pks, err := spk.QueryPublicKeyScope().All(ctx)

	if err != nil {
		return nil, fosite.ErrNotFound

	}

	return pks, nil
}

func getPksFromSpk(spk *ent.SubjectPublicKeys, keyId string, ctx context.Context) (*ent.PublicKeyScopes, error) {
	pks, err := spk.QueryPublicKeyScope().Where(publickeyscopes.ID(keyId)).Only(ctx)
	if err != nil {
		return nil, fosite.ErrNotFound

	}
	return pks, nil
}

func getSpkFromIpk(isp *ent.IssuerPublicKeys, subject string, ctx context.Context) (*ent.SubjectPublicKeys, error) {
	spk, err := isp.QuerySubjectPublicKey().Where(subjectpublickeys.ID(subject)).Only(ctx)
	if err != nil {
		return nil, fosite.ErrNotFound

	}
	return spk, nil
}

func getIPK(client *ent.Client, ctx context.Context, issuer string) (*ent.IssuerPublicKeys, error) {
	isp, err := client.IssuerPublicKeys.Get(ctx, issuer)
	if err != nil {
		return nil, fosite.ErrNotFound
	}
	return isp, nil
}
