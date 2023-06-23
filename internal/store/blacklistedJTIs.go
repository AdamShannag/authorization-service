package store

import (
	"authorization-service/ent"
	"authorization-service/ent/blacklistedjtis"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ory/fosite"
)

func createBlacklistedJTI(ctx context.Context, client *ent.Client, jit string, exp time.Time) error {
	if blacklistedJTIExists(ctx, client, jit) {
		return fosite.ErrJTIKnown
	}

	u, err := client.BlacklistedJTIs.
		Create().
		SetID(jit).
		SetExpiry(exp).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating a blacklistedJTI: %w", err)
	}
	log.Println("blacklistedJTI was created: ", u)
	return nil
}

func blacklistedJTIExists(ctx context.Context, client *ent.Client, jit string) bool {
	if _, err := client.BlacklistedJTIs.Get(ctx, jit); err != nil {
		return false
	}
	return true
}

func validateJWT(ctx context.Context, client *ent.Client, jti string) error {
	if bJti, err := client.BlacklistedJTIs.Query().
		Where(blacklistedjtis.ID(jti)).
		Only(ctx); err != nil && bJti.Expiry.After(time.Now()) {
		return fosite.ErrJTIKnown
	}
	return nil
}

func deleteExpiredBlacklistedJTIs(ctx context.Context, client *ent.Client) (int, error) {
	return client.BlacklistedJTIs.Delete().Where(blacklistedjtis.ExpiryLT(time.Now())).Exec(ctx)
}
