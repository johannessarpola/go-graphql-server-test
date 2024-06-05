package resolvers

import (
	"context"
	"fmt"
	"github.com/johannessarpola/graphql-server-test/graph/model"
	"github.com/johannessarpola/graphql-server-test/internal/app"
	"github.com/johannessarpola/graphql-server-test/pkg/spotify"
	"github.com/johannessarpola/graphql-server-test/pkg/transform"
)

func GetPlaylist(ctx context.Context, obj *model.AddItemsToPlaylistPayload) (*model.Playlist, error) {
	appCtx := app.GetAppContext(ctx)

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

func AddItemsToPlaylist(ctx context.Context, input model.AddItemsToPlaylistInput) (*model.AddItemsToPlaylistPayload, error) {
	appCtx := app.GetAppContext(ctx)

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

func Tracks(ctx context.Context, obj *model.Playlist) ([]*model.Track, error) {
	appCtx := app.GetAppContext(ctx)
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

func FeaturedPlaylists(ctx context.Context) ([]*model.Playlist, error) {
	appCtx := app.GetAppContext(ctx)
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

func Playlist(ctx context.Context, id string) (*model.Playlist, error) {
	appCtx := app.GetAppContext(ctx)
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
