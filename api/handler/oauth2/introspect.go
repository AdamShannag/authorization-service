package oauth2

import (
	"authorization-service/internal/provider"
	"log"
	"net/http"
)

func (h Oauth2) Introspect(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	mySessionData := provider.NewSession("")
	ir, err := h.oauth2.NewIntrospectionRequest(ctx, req, mySessionData)
	if err != nil {
		log.Printf("Error occurred in NewIntrospectionRequest: %+v", err)
		h.oauth2.WriteIntrospectionError(ctx, w, err)
		return
	}

	h.oauth2.WriteIntrospectionResponse(ctx, w, ir)
}
