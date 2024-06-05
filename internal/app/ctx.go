package app

import (
	"context"
	"github.com/johannessarpola/go-graphql-server-test/pkg/auth"
	"github.com/johannessarpola/go-graphql-server-test/pkg/spotify"
)

type CustomContext struct {
	SpotifyAPI  *spotify.API
	UserDetails auth.UserDetails
}

const appContextKey = "app"

func GetAppContext(ctx context.Context) *CustomContext {
	appCtx, ok := ctx.Value(appContextKey).(*CustomContext)
	if !ok {
		return nil
	}
	return appCtx
}

func WithAppContext(ctx context.Context, appCtx *CustomContext) context.Context {
	return context.WithValue(ctx, appContextKey, appCtx)
}
