package common

import (
	"context"
	"github.com/johannessarpola/graphql-test/pkg/spotify"
)

type AppContext struct {
	SpotifyAPI  *spotify.API
	UserDetails UserDetails
}

const stateLength = 32
const oauthStateContextKey = "state"
const appContextKey = "app"

func GetOauthState(ctx context.Context) string {
	oauthState, _ := ctx.Value(oauthStateContextKey).(string)
	return oauthState
}

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
