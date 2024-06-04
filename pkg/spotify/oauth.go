package spotify

import (
	"golang.org/x/oauth2"
)

func NewSpotifyOauthConfig(config AuthConfig) oauth2.Config {
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
