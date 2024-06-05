package transform

import (
	"github.com/johannessarpola/graphql-server-test/graph/model"
	"github.com/johannessarpola/graphql-server-test/pkg/spotify"
)

func Track(t spotify.Track) model.Track {
	return model.Track{
		ID:         t.Id,
		Name:       t.Name,
		DurationMs: t.DurationMs,
		Explicit:   t.Explicit,
		URI:        t.Uri,
	}
}

func Playlist(pl spotify.Playlist) model.Playlist {
	return model.Playlist{
		ID:          pl.Id,
		Name:        pl.Name,
		Description: &pl.Description,
	}
}
