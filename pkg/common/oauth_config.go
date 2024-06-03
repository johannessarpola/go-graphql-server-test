package common

import (
	"context"
	"fmt"
	"github.com/johannessarpola/graphql-test/pkg/spotify"
	"golang.org/x/oauth2"
)

func NewSpotifyOauthConfig(config spotify.AuthConfig) oauth2.Config {
	return oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  config.RedirectURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  config.AuthEndpoint,
			TokenURL: config.TokenEndpoint,
		},
		Scopes: config.Scopes,
	}
}

func OauthStateContext(ctx context.Context) context.Context {
	ei, ok := ctx.Value(oauthStateContextKey).(string)
	if ok {
		fmt.Printf("existing random state: %s\n", ei)
		return ctx
	} else {
		rs := GenerateRandomString(stateLength)
		fmt.Printf("new random state: %s\n", rs)
		return context.WithValue(ctx, oauthStateContextKey, rs)
	}
}
