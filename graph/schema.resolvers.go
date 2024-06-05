package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"

	"github.com/johannessarpola/go-graphql-server-test/graph/model"
	"github.com/johannessarpola/go-graphql-server-test/internal/resolvers"
)

// Playlist is the resolver for the playlist field.
func (r *addItemsToPlaylistPayloadResolver) Playlist(ctx context.Context, obj *model.AddItemsToPlaylistPayload) (*model.Playlist, error) {
	return resolvers.GetPlaylist(ctx, obj)
}

// AddItemsToPlaylist is the resolver for the addItemsToPlaylist field.
func (r *mutationResolver) AddItemsToPlaylist(ctx context.Context, input model.AddItemsToPlaylistInput) (*model.AddItemsToPlaylistPayload, error) {
	return resolvers.AddItemsToPlaylist(ctx, input)
}

// Tracks is the resolver for the tracks field.
func (r *playlistResolver) Tracks(ctx context.Context, obj *model.Playlist) ([]*model.Track, error) {
	return resolvers.Tracks(ctx, obj)
}

// FeaturedPlaylists is the resolver for the featuredPlaylists field.
func (r *queryResolver) FeaturedPlaylists(ctx context.Context) ([]*model.Playlist, error) {
	return resolvers.FeaturedPlaylists(ctx)
}

// Playlist is the resolver for the playlist field.
func (r *queryResolver) Playlist(ctx context.Context, id string) (*model.Playlist, error) {
	return resolvers.Playlist(ctx, id)
}

// AddItemsToPlaylistPayload returns AddItemsToPlaylistPayloadResolver implementation.
func (r *Resolver) AddItemsToPlaylistPayload() AddItemsToPlaylistPayloadResolver {
	return &addItemsToPlaylistPayloadResolver{r}
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Playlist returns PlaylistResolver implementation.
func (r *Resolver) Playlist() PlaylistResolver { return &playlistResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type addItemsToPlaylistPayloadResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type playlistResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
