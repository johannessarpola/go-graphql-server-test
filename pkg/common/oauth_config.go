package common

import (
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
