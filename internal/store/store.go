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

	clientsMutex                sync.RWMutex
	authorizeCodesMutex         sync.RWMutex
	idSessionsMutex             sync.RWMutex
	accessTokensMutex           sync.RWMutex
	refreshTokensMutex          sync.RWMutex
	pkcesMutex                  sync.RWMutex
	usersMutex                  sync.RWMutex
	blacklistedJTIsMutex        sync.RWMutex
	accessTokenRequestIDsMutex  sync.RWMutex
	refreshTokenRequestIDsMutex sync.RWMutex
	issuerPublicKeysMutex       sync.RWMutex
}

func NewDBStore(client *ent.Client) *DBStore {
	return &DBStore{
		client:                 client,
		AccessTokenRequestIDs:  make(map[string]string),
		RefreshTokenRequestIDs: make(map[string]string),
	}
}

func (s *DBStore) CreateOpenIDConnectSession(ctx context.Context, authorizeCode string, requester fosite.Requester) error {
	log.Println("\n\n-----------------------------\n", requester.GetRequestForm(), "---------------------------")

	s.idSessionsMutex.Lock()
	defer s.idSessionsMutex.Unlock()
	log.Printf("\n\nCreateOpenIDConnectSession: \n%v\nCODE: %s", requester, authorizeCode)
	return createIDSession(ctx, s.client, authorizeCode, requester)
}

func (s *DBStore) GetOpenIDConnectSession(ctx context.Context, authorizeCode string, requester fosite.Requester) (fosite.Requester, error) {
	s.idSessionsMutex.RLock()
	defer s.idSessionsMutex.RUnlock()
	log.Printf("\n\nGetOpenIDConnectSession,\nCODE: %s", authorizeCode)
	r, err := findIDSessionByCode(ctx, s.client, authorizeCode)

	log.Printf("\n\nGetOpenIDConnectSession,\nIDSESSION: %v", r)

	return r, err
}

// DeleteOpenIDConnectSession is not really called from anywhere and it is deprecated.
func (s *DBStore) DeleteOpenIDConnectSession(ctx context.Context, authorizeCode string) error {
	s.idSessionsMutex.Lock()
	defer s.idSessionsMutex.Unlock()
	log.Printf("\n\nDeleteOpenIDConnectSession: %v", authorizeCode)
	return deleteIDSessionByCode(ctx, s.client, authorizeCode)
}

func (s *DBStore) GetClient(ctx context.Context, id string) (fosite.Client, error) {
	s.clientsMutex.RLock()
	defer s.clientsMutex.RUnlock()
	log.Printf("\n\nGetClient: %v", id)
	return getClient(ctx, s.client, id)
}

func (s *DBStore) ClientAssertionJWTValid(ctx context.Context, jti string) error {
	s.blacklistedJTIsMutex.RLock()
	defer s.blacklistedJTIsMutex.RUnlock()
	log.Printf("\n\nClientAssertionJWTValid: %v", jti)
	return validateJWT(ctx, s.client, jti)
}

func (s *DBStore) SetClientAssertionJWT(ctx context.Context, jti string, exp time.Time) error {
	s.blacklistedJTIsMutex.Lock()
	defer s.blacklistedJTIsMutex.Unlock()
	log.Printf("\n\nSetClientAssertionJWT: %v", jti)

	_, err := deleteExpiredBlacklistedJTIs(ctx, s.client)

	if err != nil {
		return err
	}

	return createBlacklistedJTI(ctx, s.client, jti, exp)
}

func (s *DBStore) CreateAuthorizeCodeSession(ctx context.Context, code string, req fosite.Requester) error {
	log.Println("\n\n-----------------------------\n", req.GetRequestForm(), "---------------------------")
	s.authorizeCodesMutex.Lock()
	defer s.authorizeCodesMutex.Unlock()
	log.Printf("\n\nCreateAuthorizeCodeSession: \n\nCODE %v", code)
	return createAuthorizeCode(ctx, s.client, code, req)
}

func (s *DBStore) GetAuthorizeCodeSession(ctx context.Context, code string, se fosite.Session) (fosite.Requester, error) {
	s.authorizeCodesMutex.RLock()
	defer s.authorizeCodesMutex.RUnlock()
	log.Printf("\n\nGetAuthorizeCodeSession: \nCODE: %v\n\n%v\n", code, se)

	rel, ok, err := findAuthorizeCodeByID(ctx, s.client, code)
	if err != nil {
		return nil, fosite.ErrNotFound
	}
	if !ok {
		return rel, fosite.ErrInvalidatedAuthorizeCode
	}
	return rel, nil
}

func (s *DBStore) InvalidateAuthorizeCodeSession(ctx context.Context, code string) error {
	s.authorizeCodesMutex.Lock()
	defer s.authorizeCodesMutex.Unlock()
	log.Printf("\n\nInvalidateAuthorizeCodeSession: \n\nCODE: %v", code)
	return updateAuthorizeCodeStatusByID(ctx, s.client, code, false)
}

func (s *DBStore) CreatePKCERequestSession(ctx context.Context, code string, req fosite.Requester) error {
	s.pkcesMutex.Lock()
	defer s.pkcesMutex.Unlock()
	log.Printf("\n\nCreatePKCERequestSession: \n\nCODE: %v", code)
	return createPKCE(ctx, s.client, code, req)
}

func (s *DBStore) GetPKCERequestSession(ctx context.Context, code string, _ fosite.Session) (fosite.Requester, error) {
	s.pkcesMutex.RLock()
	defer s.pkcesMutex.RUnlock()
	log.Printf("\n\nGetPKCERequestSession: %v", code)
	return findPKCEByCode(ctx, s.client, code)
}

func (s *DBStore) DeletePKCERequestSession(ctx context.Context, code string) error {
	s.pkcesMutex.Lock()
	defer s.pkcesMutex.Unlock()
	log.Printf("\n\nDeletePKCERequestSession: %v", code)
	return deletePKCEByCode(ctx, s.client, code)
}

func (s *DBStore) CreateAccessTokenSession(ctx context.Context, signature string, req fosite.Requester) error {
	s.accessTokenRequestIDsMutex.Lock()
	defer s.accessTokenRequestIDsMutex.Unlock()
	s.accessTokensMutex.Lock()
	defer s.accessTokensMutex.Unlock()
	log.Printf("\n\nCreateAccessTokenSession: \n\nsignature: %v", signature)
	err := createAccessToken(ctx, s.client, signature, req)
	s.AccessTokenRequestIDs[req.GetID()] = signature
	return err
}

func (s *DBStore) GetAccessTokenSession(ctx context.Context, signature string, _ fosite.Session) (fosite.Requester, error) {
	s.accessTokensMutex.RLock()
	defer s.accessTokensMutex.RUnlock()
	log.Printf("\n\nGetAccessTokenSession: \n\nsignature: %v", signature)
	return findAccessTokenBySignature(ctx, s.client, signature)
}

func (s *DBStore) DeleteAccessTokenSession(ctx context.Context, signature string) error {
	s.accessTokensMutex.Lock()
	defer s.accessTokensMutex.Unlock()
	log.Printf("\n\nDeleteAccessTokenSession: %v", signature)
	return deleteAccessTokenBySignature(ctx, s.client, signature)
}

func (s *DBStore) CreateRefreshTokenSession(ctx context.Context, signature string, req fosite.Requester) error {
	s.refreshTokenRequestIDsMutex.Lock()
	defer s.refreshTokenRequestIDsMutex.Unlock()
	s.refreshTokensMutex.Lock()
	defer s.refreshTokensMutex.Unlock()
	log.Printf("\n\nCreateRefreshTokenSession: \n\nsignature: %v", signature)
	err := createRefreshToken(ctx, s.client, signature, req)
	if err != nil {
		return err
	}
	s.RefreshTokenRequestIDs[req.GetID()] = signature
	return err
}

func (s *DBStore) GetRefreshTokenSession(ctx context.Context, signature string, _ fosite.Session) (fosite.Requester, error) {
	s.refreshTokensMutex.RLock()
	defer s.refreshTokensMutex.RUnlock()
	log.Printf("\n\nGetRefreshTokenSession: \n\nsignature: %v", signature)
	rel, ok, err := findRefreshTokenByID(ctx, s.client, signature)
	if err != nil {
		return nil, fosite.ErrNotFound
	}
	if !ok {
		return rel, fosite.ErrInactiveToken
	}
	return rel, nil
}

func (s *DBStore) DeleteRefreshTokenSession(ctx context.Context, signature string) error {
	s.refreshTokensMutex.Lock()
	defer s.refreshTokensMutex.Unlock()
	log.Printf("\n\nDeleteRefreshTokenSession: %v", signature)
	return deleteRefreshTokenByID(ctx, s.client, signature)
}

func (s *DBStore) Authenticate(ctx context.Context, name string, secret string) error {
	s.usersMutex.RLock()
	defer s.usersMutex.RUnlock()
	log.Printf("\n\nAuthenticate: %v", name)
	return authenticateUser(ctx, s.client, name, secret)
}

func (s *DBStore) RevokeRefreshToken(ctx context.Context, requestID string) error {
	s.refreshTokenRequestIDsMutex.Lock()
	defer s.refreshTokenRequestIDsMutex.Unlock()
	log.Printf("\n\nRevokeRefreshToken: %v", requestID)
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
	log.Printf("\n\nRevokeAccessToken: %v", requestID)
	if signature, exists := s.AccessTokenRequestIDs[requestID]; exists {
		if err := s.DeleteAccessTokenSession(ctx, signature); err != nil {
			return err
		}
	}
	return nil
}

func (s *DBStore) GetPublicKey(ctx context.Context, issuer string, subject string, keyId string) (*jose.JSONWebKey, error) {
	s.issuerPublicKeysMutex.RLock()
	defer s.issuerPublicKeysMutex.RUnlock()
	// log.Printf("\n\nGetPublicKey: %v", issuer)
	return getPublicKey(ctx, s.client, issuer, subject, keyId)
}

func (s *DBStore) GetPublicKeys(ctx context.Context, issuer string, subject string) (*jose.JSONWebKeySet, error) {
	s.issuerPublicKeysMutex.RLock()
	defer s.issuerPublicKeysMutex.RUnlock()
	// log.Printf("\n\nGetPublicKeys: %v", issuer)
	return getPublicKeys(ctx, s.client, issuer, subject)
}

func (s *DBStore) GetPublicKeyScopes(ctx context.Context, issuer string, subject string, keyId string) ([]string, error) {
	s.issuerPublicKeysMutex.RLock()
	defer s.issuerPublicKeysMutex.RUnlock()
	// log.Printf("\n\nGetPublicKeyScopes: %v", issuer)
	return getPublicKeyScopes(ctx, s.client, issuer, subject, keyId)
}

func (s *DBStore) IsJWTUsed(ctx context.Context, jti string) (bool, error) {
	log.Printf("\n\nIsJWTUsed: %v", jti)
	err := s.ClientAssertionJWTValid(ctx, jti)
	if err != nil {
		return true, nil
	}

	return false, nil
}

func (s *DBStore) MarkJWTUsedForTime(ctx context.Context, jti string, exp time.Time) error {
	log.Printf("\n\nMarkJWTUsedForTime: %v", jti)
	return s.SetClientAssertionJWT(ctx, jti, exp)
}
