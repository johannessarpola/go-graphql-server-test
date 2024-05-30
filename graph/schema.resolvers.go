package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"fmt"
	"github.com/johannessarpola/graphql-test/pkg/transform"

	"github.com/johannessarpola/graphql-test/graph/model"
	"github.com/johannessarpola/graphql-test/pkg/common"
	"github.com/johannessarpola/graphql-test/pkg/spotify"
)

// Playlist is the resolver for the playlist field.
func (r *addItemsToPlaylistPayloadResolver) Playlist(ctx context.Context, obj *model.AddItemsToPlaylistPayload) (*model.Playlist, error) {
	appCtx := common.GetContext(ctx)

	if obj.Playlist != nil {
		pl, err := appCtx.SpotifyAPI.GetPlaylist(obj.Playlist.ID)
		if err != nil {
			return nil, err
		}
		tt := transform.Playlist(*pl)
		return &tt, nil
	} else {
		return nil, fmt.Errorf("playlist not found")
	}
}

// AddItemsToPlaylist is the resolver for the addItemsToPlaylist field.
func (r *mutationResolver) AddItemsToPlaylist(ctx context.Context, input model.AddItemsToPlaylistInput) (*model.AddItemsToPlaylistPayload, error) {
	appCtx := common.GetContext(ctx)

	pd := spotify.AddTracksPayload{
		Id:   input.PlaylistID,
		Uris: input.Uris,
	}
	rs, err := appCtx.SpotifyAPI.AddTrackToPlaylist(pd)
	if err != nil {
		return nil, err
	}

	if len(rs.SnapshotId) > 0 {
		return &model.AddItemsToPlaylistPayload{
			Code:    200,
			Success: true,
			Message: "Tracks added successfully",
			Playlist: &model.Playlist{
				ID: input.PlaylistID,
			},
		}, nil
	} else {
		return &model.AddItemsToPlaylistPayload{
			Code:     500,
			Success:  false,
			Message:  rs.Error, // Could be fancier
			Playlist: nil,
		}, nil
	}
}

// Tracks is the resolver for the tracks field.
func (r *playlistResolver) Tracks(ctx context.Context, obj *model.Playlist) ([]*model.Track, error) {
	appCtx := common.GetContext(ctx)
	fmt.Println("resolver.Tracks")
	var l []*model.Track
	ts, err := appCtx.SpotifyAPI.GetTracks(obj.ID)

	if err != nil {
		return nil, err
	}

	for _, t := range ts {
		tt := transform.Track(t)
		l = append(l, &tt)
	}
	return l, nil
}

// FeaturedPlaylists is the resolver for the featuredPlaylists field.
func (r *queryResolver) FeaturedPlaylists(ctx context.Context) ([]*model.Playlist, error) {
	appCtx := common.GetContext(ctx)
	apiData, err := appCtx.SpotifyAPI.GetFeaturedPlaylists()
	if err != nil {
		return nil, err
	}

	var playlists []*model.Playlist
	for _, pl := range apiData.Playlists.Items {
		pp := transform.Playlist(pl)
		playlists = append(playlists, &pp)
	}

	return playlists, nil
}

// Playlist is the resolver for the playlist field.
func (r *queryResolver) Playlist(ctx context.Context, id string) (*model.Playlist, error) {
	appCtx := common.GetContext(ctx)
	fmt.Println("resolver.Playlist")
	rs, err := appCtx.SpotifyAPI.GetPlaylist(id)
	if err != nil {
		return nil, err
	}

	return &model.Playlist{
		ID:          rs.Id,
		Name:        rs.Name,
		Description: &rs.Description,
	}, nil
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
