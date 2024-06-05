package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-resty/resty/v2"
	"github.com/johannessarpola/go-graphql-server-test/graph"
	"github.com/johannessarpola/go-graphql-server-test/internal/app"
	"github.com/johannessarpola/go-graphql-server-test/pkg/auth"
	"github.com/johannessarpola/go-graphql-server-test/pkg/spotify"
	"github.com/johannessarpola/go-graphql-server-test/pkg/state"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

var appConfig app.Config
var oauthConfig oauth2.Config
var stateCache state.Cache

func init() {
	stateCache = state.NewStateCache()
	var err error
	appConfig, err = app.LoadConfig[app.Config]("config/config.dev.yaml")
	if err != nil {
		fmt.Println("Error loading config from yaml")
		panic(err)
	}
	oauthConfig = spotify.NewSpotifyOauthConfig(appConfig.SpotifyConfig.Auth)
}

func main() {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	appCtx := &app.CustomContext{
		// This is still unauthenticated client
		SpotifyAPI: spotify.NewSpotifyAPI(appConfig.SpotifyConfig.Base, resty.New()),
		UserDetails: auth.UserDetails{
			Login:         "not_logged_in",
			Authenticated: false,
		},
	}

	http.Handle("/", http.HandlerFunc(handleHome))
	http.Handle("/login", http.HandlerFunc(handleLogin))
	http.Handle("/callback", withAppContext(appCtx, http.HandlerFunc(handleCallback)))
	http.Handle("/playground", withAppContext(appCtx, hasAuthentication(playground.Handler("GraphQL playground", "/query"))))
	http.Handle("/query", withAppContext(appCtx, hasAuthentication(srv)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", appConfig.Port)
	err := http.ListenAndServe(":"+appConfig.Port, nil)
	if err != nil {
		fmt.Println("failed to start server:", err)
		panic(err)
	}
}

func withAppContext(appContext *app.CustomContext, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := app.WithAppContext(r.Context(), appContext)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func hasAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app := app.GetAppContext(r.Context())
		if app == nil || !app.UserDetails.Authenticated {
			handleHome(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	var htmlIndex = `<html><body><a href="/login">Log in with Spotify</a></body></html>`
	fmt.Fprintf(w, htmlIndex)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	newState := state.NewState()
	stateCache.Add(newState, r.RemoteAddr)

	url := oauthConfig.AuthCodeURL(newState, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	appCtx := app.GetAppContext(ctx)
	inputState := r.FormValue("state")
	if !stateCache.Has(inputState) {
		http.Error(w, "Unknown state", http.StatusBadRequest)
		return
	}

	code := r.FormValue("code")
	client, err := spotify.NewAuthenticatedClient(code, &oauthConfig, ctx)
	if err != nil {
		fmt.Println("Error creating hasAuthentication client", err)
		panic(err)
	}

	appCtx.UserDetails = auth.UserDetails{
		Login:         "test@test.fi",
		Authenticated: true,
	}
	appCtx.SpotifyAPI = spotify.NewSpotifyAPI(appConfig.SpotifyConfig.Base, client)

	http.Redirect(w, r, "/playground", http.StatusPermanentRedirect)
}
