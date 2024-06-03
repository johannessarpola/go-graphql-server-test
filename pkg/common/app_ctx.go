package common

import (
	"context"
	"github.com/johannessarpola/graphql-test/pkg/spotify"
	"net/http"
)

type AppContext struct {
	SpotifyAPI  *spotify.API
	UserDetails UserDetails
}

const stateLength = 32
const oauthStateContextKey = "state"
const appContextKey = "app"

func CreateContext(args *AppContext, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		appCtx := &AppContext{
			SpotifyAPI: args.SpotifyAPI,
		}
		requestWithCtx := r.WithContext(context.WithValue(r.Context(), appContextKey, appCtx))
		next.ServeHTTP(w, requestWithCtx)
	})
}

func GetOauthState(ctx context.Context) string {
	oauthState, ok := ctx.Value(oauthStateContextKey).(string)
	if !ok {
		// For now return random string (should fail auth)
		return GenerateRandomString(stateLength)
	}
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
