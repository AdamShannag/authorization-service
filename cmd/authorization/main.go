package main

import (
	"authorization-service/api"
	"authorization-service/ent"
	"authorization-service/ent/clients"
	"authorization-service/ent/user"
	"authorization-service/internal/provider"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const webPort = "8000"

func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func main() {
	client := Open("host=localhost port=5432 user=postgres password=password dbname=fosite sslmode=disable timezone=UTC connect_timeout=5")

	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	ctx := context.Background()

	if err := seed(ctx, client); err != nil {
		log.Fatalf("failed seeding the database: %v", err)
	}

	var (
		mux = api.NewMux(
			provider.NewOauth2Provider(client),
		)
		server = http.Server{
			Addr:    ":" + webPort,
			Handler: mux,
		}
	)

	log.Println("server starting: http://localhost" + server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("server error", err)
	}
}

func seed(ctx context.Context, client *ent.Client) error {
	_, err := client.User.Query().
		Where(
			user.ID("peter"),
		).
		Only(ctx)

	switch {
	case ent.IsNotFound(err):
		_, err = client.User.Create().
			SetID("peter").
			SetPassword("secret").
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed creating user: %v", err)
		}
	case err != nil:
		return fmt.Errorf("failed querying user: %v", err)
	}

	_, err = client.Clients.Query().
		Where(clients.ID("my-client")).Only(ctx)

	switch {
	case ent.IsNotFound(err):
		_, err := client.Clients.Create().
			SetID("my-client").
			SetClientSecret([]byte(`$2a$10$IxMdI6d.LIRZPpSfEwNoeu4rY3FhDREsxFJXikcgdRRAStxUlsuEO`)).
			SetRotatedSecrets([][]byte{[]byte(`$2y$10$X51gLxUQJ.hGw1epgHTE5u0bt64xM0COU7K9iAp.OFg8p2pUd.1zC `)}).
			SetRedirectUris([]string{"http://localhost:3846/callback"}).
			SetResponseTypes([]string{"id_token", "code", "token", "id_token token", "code id_token", "code token", "code id_token token"}).
			SetGrantTypes([]string{"implicit", "refresh_token", "authorization_code", "password", "client_credentials"}).
			SetScopes([]string{"fosite", "openid", "photos", "offline"}).
			SetPublic(true).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed creating client: %v", err)
		}
	case err != nil:
		return fmt.Errorf("failed querying client: %v", err)
	}

	_, err = client.Clients.Query().
		Where(clients.ID("encoded:client")).Only(ctx)

	switch {
	case ent.IsNotFound(err):
		_, err := client.Clients.Create().
			SetID("encoded:client").
			SetClientSecret([]byte(`$2a$10$A7M8b65dSSKGHF0H2sNkn.9Z0hT8U1Nv6OWPV3teUUaczXkVkxuDS`)).
			SetRotatedSecrets(nil).
			SetRedirectUris([]string{"http://localhost:3846/callback"}).
			SetResponseTypes([]string{"id_token", "code", "token", "id_token token", "code id_token", "code token", "code id_token token"}).
			SetGrantTypes([]string{"implicit", "refresh_token", "authorization_code", "password", "client_credentials"}).
			SetScopes([]string{"fosite", "openid", "photos", "offline"}).
			SetPublic(true).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed creating client: %v", err)
		}
	case err != nil:
		return fmt.Errorf("failed querying client: %v", err)
	}

	return nil
}
