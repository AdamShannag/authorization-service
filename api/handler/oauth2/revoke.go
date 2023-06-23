package oauth2

import (
	"net/http"
)

func (h Oauth2) Revoke(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	err := h.oauth2.NewRevocationRequest(ctx, req)

	h.oauth2.WriteRevocationResponse(ctx, w, err)
}
