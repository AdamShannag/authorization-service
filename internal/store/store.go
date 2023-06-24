package store

import (
	"authorization-service/ent"
	"context"
	"log"
	"sync"
	"time"

	"gopkg.in/square/go-jose.v2"

	"github.com/ory/fosite"
)

type DBStore struct {
	client *ent.Client
	// In-memory request ID to token signatures
	AccessTokenRequestIDs  map[string]string
	RefreshTokenRequestIDs map[string]string

	accessTokenRequestIDsMutex  sync.RWMutex
	refreshTokenRequestIDsMutex sync.RWMutex
}

func NewDBStore(client *ent.Client) *DBStore {
	return &DBStore{
		client:                 client,
		AccessTokenRequestIDs:  make(map[string]string),
		RefreshTokenRequestIDs: make(map[string]string),
	}
}

func (s *DBStore) CreateOpenIDConnectSession(ctx context.Context, authorizeCode string, requester fosite.Requester) error {
	// log.Printf("\n\nCreateOpenIDConnectSession: \n%v\nCODE: %s", requester, authorizeCode)
	err := createIDSession(ctx, s.client, authorizeCode, requester)

	return err
}

func (s *DBStore) GetOpenIDConnectSession(ctx context.Context, authorizeCode string, requester fosite.Requester) (fosite.Requester, error) {
	r, err := findIDSessionByCode(ctx, s.client, authorizeCode, requester)

	// log.Printf("\n\nGetOpenIDConnectSession,\nREQUEST: %v\n", requester.GetSession())
	// log.Printf("\n\nGetOpenIDConnectSession,\nIDSESSION: %v\n", r.GetSession())

	return r, err
}

// DeleteOpenIDConnectSession is not really called from anywhere and it is deprecated.
func (s *DBStore) DeleteOpenIDConnectSession(ctx context.Context, authorizeCode string) error {
	// log.Printf("\n\nDeleteOpenIDConnectSession: %v", authorizeCode)
	return deleteIDSessionByCode(ctx, s.client, authorizeCode)
}

func (s *DBStore) GetClient(ctx context.Context, id string) (fosite.Client, error) {
	// log.Printf("\n\nGetClient: %v", id)
	return getClient(ctx, s.client, id)
}

func (s *DBStore) ClientAssertionJWTValid(ctx context.Context, jti string) error {
	// log.Printf("\n\nClientAssertionJWTValid: %v", jti)
	return validateJWT(ctx, s.client, jti)
}

func (s *DBStore) SetClientAssertionJWT(ctx context.Context, jti string, exp time.Time) error {

	// log.Printf("\n\nSetClientAssertionJWT: %v", jti)

	_, err := deleteExpiredBlacklistedJTIs(ctx, s.client)

	if err != nil {
		return err
	}

	return createBlacklistedJTI(ctx, s.client, jti, exp)
}

func (s *DBStore) CreateAuthorizeCodeSession(ctx context.Context, code string, req fosite.Requester) error {
	// log.Printf("\n\nCreateAuthorizeCodeSession: \n\nCODE %v", code)
	return createAuthorizeCode(ctx, s.client, code, req)
}

func (s *DBStore) GetAuthorizeCodeSession(ctx context.Context, code string, se fosite.Session) (fosite.Requester, error) {
	log.Printf("\n==========================================\n\nGetAuthorizeCodeSession: %+v\n\n===============================\n\n", se)

	rel, ok, err := findAuthorizeCodeByID(ctx, s.client, code)
	if err != nil {
		return nil, fosite.ErrNotFound
	}
	if !ok {
		return rel, fosite.ErrInvalidatedAuthorizeCode
	}
	rel.SetSession(se)
	return rel, nil
}

func (s *DBStore) InvalidateAuthorizeCodeSession(ctx context.Context, code string) error {
	// log.Printf("\n\nInvalidateAuthorizeCodeSession: \n\nCODE: %v", code)
	return updateAuthorizeCodeStatusByID(ctx, s.client, code, false)
}

func (s *DBStore) CreatePKCERequestSession(ctx context.Context, code string, req fosite.Requester) error {
	// log.Printf("\n\nCreatePKCERequestSession: \n\nCODE: %v", code)
	return createPKCE(ctx, s.client, code, req)
}

func (s *DBStore) GetPKCERequestSession(ctx context.Context, code string, se fosite.Session) (fosite.Requester, error) {
	// log.Printf("\n\nGetPKCERequestSession: %v", code)
	rel, err := findPKCEByCode(ctx, s.client, code)

	rel.SetSession(se)

	return rel, err

}

func (s *DBStore) DeletePKCERequestSession(ctx context.Context, code string) error {
	// log.Printf("\n\nDeletePKCERequestSession: %v", code)
	return deletePKCEByCode(ctx, s.client, code)
}

func (s *DBStore) CreateAccessTokenSession(ctx context.Context, signature string, req fosite.Requester) error {
	s.accessTokenRequestIDsMutex.Lock()
	defer s.accessTokenRequestIDsMutex.Unlock()
	// log.Printf("\n\nCreateAccessTokenSession: \n\nsignature: %v", signature)
	err := createAccessToken(ctx, s.client, signature, req)
	s.AccessTokenRequestIDs[req.GetID()] = signature
	return err
}

func (s *DBStore) GetAccessTokenSession(ctx context.Context, signature string, se fosite.Session) (fosite.Requester, error) {
	// log.Printf("\n\nGetAccessTokenSession: \n\nsignature: %v", signature)
	r, err := findAccessTokenBySignature(ctx, s.client, signature)
	r.SetSession(se)
	return r, err
}

func (s *DBStore) DeleteAccessTokenSession(ctx context.Context, signature string) error {
	// log.Printf("\n\nDeleteAccessTokenSession: %v", signature)
	return deleteAccessTokenBySignature(ctx, s.client, signature)
}

func (s *DBStore) CreateRefreshTokenSession(ctx context.Context, signature string, req fosite.Requester) error {
	s.refreshTokenRequestIDsMutex.Lock()
	defer s.refreshTokenRequestIDsMutex.Unlock()
	// log.Printf("\n\nCreateRefreshTokenSession: \n\nsignature: %v", signature)
	err := createRefreshToken(ctx, s.client, signature, req)
	if err != nil {
		return err
	}
	s.RefreshTokenRequestIDs[req.GetID()] = signature
	return err
}

func (s *DBStore) GetRefreshTokenSession(ctx context.Context, signature string, se fosite.Session) (fosite.Requester, error) {
	// log.Printf("\n\nGetRefreshTokenSession: \n\nsignature: %v", signature)
	rel, ok, err := findRefreshTokenByID(ctx, s.client, signature)
	if err != nil {
		return nil, fosite.ErrNotFound
	}
	if !ok {
		return rel, fosite.ErrInactiveToken
	}
	rel.SetSession(se)
	return rel, nil
}

func (s *DBStore) DeleteRefreshTokenSession(ctx context.Context, signature string) error {
	// log.Printf("\n\nDeleteRefreshTokenSession: %v", signature)
	return deleteRefreshTokenByID(ctx, s.client, signature)
}

func (s *DBStore) Authenticate(ctx context.Context, name string, secret string) error {
	// log.Printf("\n\nAuthenticate: %v", name)
	return authenticateUser(ctx, s.client, name, secret)
}

func (s *DBStore) RevokeRefreshToken(ctx context.Context, requestID string) error {
	s.refreshTokenRequestIDsMutex.Lock()
	defer s.refreshTokenRequestIDsMutex.Unlock()
	// log.Printf("\n\nRevokeRefreshToken: %v", requestID)
	if signature, exists := s.RefreshTokenRequestIDs[requestID]; exists {
		if err := updateRefreshTokenStatusByID(ctx, s.client, signature, false); err != nil {
			return err
		}
	}
	return nil
}

func (s *DBStore) RevokeRefreshTokenMaybeGracePeriod(ctx context.Context, requestID string, signature string) error {
	// no configuration option is available; grace period is not available with memory store
	return s.RevokeRefreshToken(ctx, requestID)
}

func (s *DBStore) RevokeAccessToken(ctx context.Context, requestID string) error {
	s.accessTokenRequestIDsMutex.RLock()
	defer s.accessTokenRequestIDsMutex.RUnlock()
	// log.Printf("\n\nRevokeAccessToken: %v", requestID)
	if signature, exists := s.AccessTokenRequestIDs[requestID]; exists {
		if err := s.DeleteAccessTokenSession(ctx, signature); err != nil {
			return err
		}
	}
	return nil
}

func (s *DBStore) GetPublicKey(ctx context.Context, issuer string, subject string, keyId string) (*jose.JSONWebKey, error) {
	return getPublicKey(ctx, s.client, issuer, subject, keyId)
}

func (s *DBStore) GetPublicKeys(ctx context.Context, issuer string, subject string) (*jose.JSONWebKeySet, error) {
	return getPublicKeys(ctx, s.client, issuer, subject)
}

func (s *DBStore) GetPublicKeyScopes(ctx context.Context, issuer string, subject string, keyId string) ([]string, error) {
	return getPublicKeyScopes(ctx, s.client, issuer, subject, keyId)
}

func (s *DBStore) IsJWTUsed(ctx context.Context, jti string) (bool, error) {
	// log.Printf("\n\nIsJWTUsed: %v", jti)
	err := s.ClientAssertionJWTValid(ctx, jti)
	if err != nil {
		return true, nil
	}

	return false, nil
}

func (s *DBStore) MarkJWTUsedForTime(ctx context.Context, jti string, exp time.Time) error {
	// log.Printf("\n\nMarkJWTUsedForTime: %v", jti)
	return s.SetClientAssertionJWT(ctx, jti, exp)
}
