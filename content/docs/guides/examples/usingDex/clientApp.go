package main

import (
    "net/http"

    "github.com/coreos/go-oidc/v3/oidc"
    "golang.org/x/oauth2"
)

func main() {

    // Initialize a provider by specifying dex's issuer URL.
    provider, err := oidc.NewProvider(ctx, "https://dex-issuer-url.com")
    if err != nil {
    // handle error
    }

    // Configure the OAuth2 config with the client values.
    oauth2Config := oauth2.Config{
        // client_id and client_secret of the client.
        ClientID:     "example-app",
        ClientSecret: "example-app-secret",

        // The redirectURL.
        RedirectURL: "http://127.0.0.1:5555/callback",

        // Discovery returns the OAuth2 endpoints.
        Endpoint: provider.Endpoint(),

        // "openid" is a required scope for OpenID Connect flows.
        //
        // Other scopes, such as "groups" can be requested.
        Scopes: []string{oidc.ScopeOpenID, "profile", "email", "groups"},
    }

}

// Create an ID token parser.
idTokenVerifier := provider.Verifier(&oidc.Config{ClientID: "example-app"})
}

// handleRedirect
//      unauthenticated users are redirected -- to -- start an OAuth2 flow + dex server
func handleRedirect(w http.ResponseWriter, r *http.Request) {
    state := newState()
    http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusFound)
}

func handleOAuth2Callback(w http.ResponseWriter, r *http.Request) {
    state := r.URL.Query().Get("state")

    // Verify state.

    oauth2Token, err := oauth2Config.Exchange(ctx, r.URL.Query().Get("code"))
    if err != nil {
        // handle error
    }

    // Extract the ID Token from OAuth2 token.
    rawIDToken, ok := oauth2Token.Extra("id_token").(string)
    if !ok {
        // handle missing token
    }

    // Parse and verify ID Token payload.
    idToken, err := idTokenVerifier.Verify(ctx, rawIDToken)
    if err != nil {
        // handle error
    }

    // Extract custom claims.
    var claims struct {
        Email    string   `json:"email"`
        Verified bool     `json:"email_verified"`
        Groups   []string `json:"groups"`
    }
    if err := idToken.Claims(&claims); err != nil {
        // handle error
    }
}
