package spotify

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log/slog"
)

type SpotifyAPI struct {
	Base   string
	client *resty.Client
}

func NewSpotifyAPI(base string) *SpotifyAPI {
	return &SpotifyAPI{
		Base:   base,
		client: resty.New(),
	}
}

func (s *SpotifyAPI) GetFeaturedPlaylists() ([]Playlist, error) {
	p := fmt.Sprintf("%s/browse/featured-playlists", s.Base)
	result := PlaylistResponse{}
	resp, err := s.client.R().
		SetResult(&result).
		Get(p)

	if resp != nil {
		slog.Info(
			"Request completed",
			"status",
			resp.Status(),
			"path",
			p,
		)
	}

	if err != nil {
		return nil, err
	}

	return result.Playlists.Items, nil
}
