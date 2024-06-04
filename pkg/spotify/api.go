package spotify

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log/slog"
)

type API struct {
	Base   string
	client *resty.Client
}

func NewSpotifyAPI(base string, client *resty.Client) *API {
	return &API{
		Base:   base,
		client: client,
	}
}

func logResponse(response *resty.Response, path string) {
	if response != nil {
		slog.Info(
			"Request completed",
			"status",
			response.Status(),
			"path",
			path,
		)
	} else {
		slog.Warn("Request was nil", "path", path)
	}
}

func (s *API) GetFeaturedPlaylists() (*GetFeaturedPlaylists, error) {
	path := fmt.Sprintf("%s/browse/featured-playlists", s.Base)
	var result GetFeaturedPlaylists
	resp, err := s.client.R().
		SetResult(&result).
		Get(path)

	logResponse(resp, path)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *API) GetPlaylist(id string) (*Playlist, error) {
	path := fmt.Sprintf("%s/playlists/%s", s.Base, id)
	var result Playlist
	resp, err := s.client.R().
		SetResult(&result).
		Get(path)

	logResponse(resp, path)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *API) GetTracks(id string) ([]Track, error) {
	path := fmt.Sprintf("%s/playlists/%s/tracks", s.Base, id)

	var result GetTracks
	resp, err := s.client.R().
		SetResult(&result).
		Get(path)

	logResponse(resp, path)
	if err != nil {
		return nil, err
	}
	var tracks []Track
	for _, item := range result.Items {
		tracks = append(tracks, item.Track)
	}

	return tracks, nil
}

func (s *API) AddTrackToPlaylist(pd AddTracksPayload) (*SnapshotOrError, error) {
	path := fmt.Sprintf("%s/playlists/%s/tracks", s.Base, pd.Id)
	var result SnapshotOrError
	reqBody := AddTracksRequest{
		Uris: pd.Uris,
	}
	resp, err := s.client.R().
		SetResult(&result).
		SetBody(reqBody).
		Post(path)

	logResponse(resp, path)
	if err != nil {
		return nil, err
	}

	return &result, nil

}
