package oauth2

import (
	"authorization-service/internal/provider"
	"log"
	"net/http"
)

func (h Oauth2) Token(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	mySessionData := provider.NewSession("")

	accessRequest, err := h.oauth2.NewAccessRequest(ctx, req, mySessionData)

	if err != nil {
		log.Printf("Error occurred in NewAccessRequest: %+v", err)
		h.oauth2.WriteAccessError(ctx, w, accessRequest, err)
		return
	}

	if accessRequest.GetGrantTypes().ExactOne("client_credentials") {
		for _, scope := range accessRequest.GetRequestedScopes() {
			accessRequest.GrantScope(scope)
		}
	}

	response, err := h.oauth2.NewAccessResponse(ctx, accessRequest)
	log.Println("()-------------------------------------->", response)
	if err != nil {
		log.Printf("Error occurred in NewAccessResponse: %+v", err)
		h.oauth2.WriteAccessError(ctx, w, accessRequest, err)
		return
	}

	h.oauth2.WriteAccessResponse(ctx, w, accessRequest, response)
}
