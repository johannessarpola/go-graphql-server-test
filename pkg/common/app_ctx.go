package common

import (
	"context"
	"github.com/johannessarpola/graphql-test/pkg/spotify"
	"net/http"
)

type AppCtx struct {
	SpotifyAPI *spotify.API
}

var appCtxKey string = "APP_CONTEXT"

func CreateContext(args *AppCtx, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		appCtx := &AppCtx{
			SpotifyAPI: args.SpotifyAPI,
		}
		requestWithCtx := r.WithContext(context.WithValue(r.Context(), appCtxKey, appCtx))
		next.ServeHTTP(w, requestWithCtx)
	})
}

func GetContext(ctx context.Context) *AppCtx {
	appCtx, ok := ctx.Value(appCtxKey).(*AppCtx)
	if !ok {
		return nil
	}
	return appCtx
}
