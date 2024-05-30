package main

import (
	"github.com/johannessarpola/graphql-test/pkg/common"
	"github.com/johannessarpola/graphql-test/pkg/spotify"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/johannessarpola/graphql-test/graph"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	spop := "https://spotify-demo-api-fe224840a08c.herokuapp.com/v1"
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	appCtx := &common.AppCtx{
		SpotifyAPI: spotify.NewSpotifyAPI(spop),
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", common.CreateContext(appCtx, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
