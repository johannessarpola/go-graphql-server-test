package spotify

import (
	"context"
	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2"
	"time"
)

func NewAuthenticatedClient(code string, oauthConfig *oauth2.Config, ctx context.Context) (*resty.Client, error) {
	token, err := oauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	client := resty.New()
	client.SetAuthToken(token.AccessToken)
	client.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
		// Refresh existingTOken if expired
		if token.Expiry.Before(time.Now()) {
			tokenSource := oauthConfig.TokenSource(ctx, token)
			newToken, err := tokenSource.Token()
			if err != nil {
				return err
			}
			token.AccessToken = newToken.AccessToken
			client.SetAuthToken(token.AccessToken)
		}
		return nil
	})
	return client, nil
}
