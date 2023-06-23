package oauth2

import (
	"authorization-service/internal/provider"
	"fmt"
	"log"
	"net/http"
)

func (h Oauth2) Auth(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ar, err := h.oauth2.NewAuthorizeRequest(ctx, req)
	if err != nil {
		log.Printf("Error occurred in NewAuthorizeRequest: %+v", err)
		h.oauth2.WriteAuthorizeError(ctx, w, ar, err)
		return
	}

	var requestedScopes string
	for _, this := range ar.GetRequestedScopes() {
		requestedScopes += fmt.Sprintf(`<li><input type="checkbox" name="scopes" value="%s">%s</li>`, this, this)
	}

	req.ParseForm()
	if req.PostForm.Get("username") != "peter" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(`<h1>Login page</h1>`))
		w.Write([]byte(fmt.Sprintf(`
				<p>Howdy! This is the log in page. For this example, it is enough to supply the username.</p>
				<form method="post">
					<p>
						By logging in, you consent to grant these scopes:
						<ul>%s</ul>
					</p>
					<input type="text" name="username" /> <small>try peter</small><br>
					<input type="submit">
				</form>
			`, requestedScopes)))
		return
	}

	for _, scope := range req.PostForm["scopes"] {
		ar.GrantScope(scope)
	}

	mySessionData := provider.NewSession("peter")

	response, err := h.oauth2.NewAuthorizeResponse(ctx, ar, mySessionData)

	if err != nil {
		log.Printf("Error occurred in NewAuthorizeResponse: %+v", err)
		h.oauth2.WriteAuthorizeError(ctx, w, ar, err)
		return
	}

	h.oauth2.WriteAuthorizeResponse(ctx, w, ar, response)
}
