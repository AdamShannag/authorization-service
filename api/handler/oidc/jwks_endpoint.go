package oidc

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"net/http"

	"gopkg.in/square/go-jose.v2"
)

func (p *Oidc) JSONWebKeysEndpoint(w http.ResponseWriter, r *http.Request) {
	set := &jose.JSONWebKeySet{
		Keys: []jose.JSONWebKey{
			{
				KeyID: "bar",
				Use:   "sig",
				Key:   &MustRSAKey().PublicKey,
			},
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(set)
}

func MustRSAKey() *rsa.PrivateKey {
	// #nosec
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	return key
}
