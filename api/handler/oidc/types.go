package oidc

type OpenIDConfiguration struct {
	Issuer                             string   `json:"issuer"`
	AuthorizationEndpoint              string   `json:"authorization_endpoint"`
	TokenEndpoint                      string   `json:"token_endpoint"`
	UserInfoEndpoint                   string   `json:"userinfo_endpoint"`
	JwksURI                            string   `json:"jwks_uri"`
	ResponseTypesSupported             []string `json:"response_types_supported"`
	SubjectTypesSupported              []string `json:"subject_types_supported"`
	IDTokenSigningAlgValuesSupported   []string `json:"id_token_signing_alg_values_supported"`
	ScopesSupported                    []string `json:"scopes_supported"`
	ClaimsSupported                    []string `json:"claims_supported"`
	GrantTypesSupported                []string `json:"grant_types_supported"`
	TokenEndpointAuthMethodsSupported  []string `json:"token_endpoint_auth_methods_supported"`
	RevocationEndpoint                 string   `json:"revocation_endpoint"`
	EndSessionEndpoint                 string   `json:"end_session_endpoint"`
	IntrospectionEndpoint              string   `json:"introspection_endpoint"`
	CheckSessionIframe                 string   `json:"check_session_iframe"`
	FrontChannelLogoutSupported        bool     `json:"frontchannel_logout_supported"`
	FrontChannelLogoutSessionSupported bool     `json:"frontchannel_logout_session_supported"`
	BackChannelLogoutSupported         bool     `json:"backchannel_logout_supported"`
	BackChannelLogoutSessionSupported  bool     `json:"backchannel_logout_session_supported"`
}
