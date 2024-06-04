package common

import (
	"context"
	"github.com/johannessarpola/graphql-test/pkg/spotify"
)

type AppContext struct {
	SpotifyAPI  *spotify.API
	UserDetails UserDetails
}

const stateLength = 64
const appContextKey = "app"

func GetAppContext(ctx context.Context) *AppContext {
	appCtx, ok := ctx.Value(appContextKey).(*AppContext)
	if !ok {
		return nil
	}
	return appCtx
}

func WithAppContext(ctx context.Context, appCtx *AppContext) context.Context {
	return context.WithValue(ctx, appContextKey, appCtx)
}
