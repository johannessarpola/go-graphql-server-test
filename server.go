package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/johannessarpola/graphql-test/pkg/common"
	"github.com/johannessarpola/graphql-test/pkg/spotify"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/johannessarpola/graphql-test/graph"
)

const defaultPort = "8080"
const appContextKey = "app"

var appConfig common.AppConfig
var oauthConfig oauth2.Config
var stateCache common.StateCache

func init() {
	stateCache = common.NewStateCache()
	var err error
	appConfig, err = common.Load[common.AppConfig]("config/config.dev.yaml")
	if err != nil {
		log.Fatal(err)
	}
	oauthConfig = common.NewSpotifyOauthConfig(appConfig.SpotifyConfig.Auth)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	cp := "config/config.dev.yaml"
	config, err := common.Load[common.AppConfig](cp) // TODO parameter
	if err != nil {
		log.Fatalf("could not load config from path %s\n", cp, err)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	appCtx := &common.AppContext{
		// This is still unauthenticated client
		SpotifyAPI: spotify.NewSpotifyAPI(config.SpotifyConfig.Base, resty.New()),
		UserDetails: common.UserDetails{
			Login:         "not_logged_in",
			Authenticated: false,
		},
	}

	http.Handle("/", http.HandlerFunc(handleHome))
	http.Handle("/login", http.HandlerFunc(handleLogin))
	http.Handle("/callback", withAppContext(appCtx, http.HandlerFunc(handleCallback)))
	http.Handle("/playground", withAppContext(appCtx, hasAuthentication(playground.Handler("GraphQL playground", "/query"))))
	http.Handle("/query", withAppContext(appCtx, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func withAppContext(appContext *common.AppContext, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := common.WithAppContext(r.Context(), appContext)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func hasAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app := common.GetAppContext(r.Context())
		if app == nil || !app.UserDetails.Authenticated {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
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
	newState := common.NewState()
	stateCache.Add(newState, r.RemoteAddr)

	url := oauthConfig.AuthCodeURL(newState, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	appCtx := common.GetAppContext(ctx)
	inputState := r.FormValue("state")
	if !stateCache.Has(inputState) {
		http.Error(w, "Unknown state", http.StatusBadRequest)
		return
	}

	code := r.FormValue("code")
	client, err := spotify.NewAuthenticatedClient(code, &oauthConfig, ctx)
	if err != nil {
		log.Fatal("Error creating hasAuthentication client", err)
	}

	appCtx.UserDetails = common.UserDetails{
		Login:         "test@test.fi",
		Authenticated: true,
	}
	appCtx.SpotifyAPI = spotify.NewSpotifyAPI(appConfig.SpotifyConfig.Base, client)

	http.Redirect(w, r, "/playground", http.StatusPermanentRedirect)
}
