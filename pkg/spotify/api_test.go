package spotify

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"net/http"
	"testing"
)

const testBase = "https://api.localhost/v1"

func mockClient() *resty.Client {
	// Create a Resty Client
	client := resty.New()
	// Get the underlying HTTP Client and set it to Mock
	httpmock.ActivateNonDefault(client.GetClient())

	return client
}

func TestGetTracks(t *testing.T) {
	client := mockClient()
	id := "test_id"
	path := fmt.Sprintf("%s/playlists/%s/tracks", testBase, id)

	httpmock.RegisterResponder("GET", path, func(request *http.Request) (*http.Response, error) {
		return httpmock.NewJsonResponse(200, httpmock.File("test_data/tracks.json"))
	})

	api := NewSpotifyAPI(testBase, client)
	r, err := api.GetTracks(id)

	if err != nil {
		t.Error(err)
	}

	if len(r) == 0 {
		t.Error("no tracks")
	}

}

func TestFeaturedPlaylists(t *testing.T) {
	client := mockClient()
	path := fmt.Sprintf("%s/browse/featured-playlists", testBase)

	httpmock.RegisterResponder("GET", path, func(request *http.Request) (*http.Response, error) {
		return httpmock.NewJsonResponse(200, httpmock.File("test_data/featured.json"))
	})

	api := NewSpotifyAPI(testBase, client)
	r, err := api.GetFeaturedPlaylists()
	if err != nil {
		t.Error(err)
	}

	if r == nil {
		t.Error("no featured playlist")
	}

	if len(r.Playlists.Items) == 0 {
		t.Error("no items")
	}
}

func TestGetPlaylist(t *testing.T) {
	client := mockClient()
	id := "test_id"
	path := fmt.Sprintf("%s/playlists/%s", testBase, id)

	httpmock.RegisterResponder("GET", path, func(request *http.Request) (*http.Response, error) {
		return httpmock.NewJsonResponse(200, httpmock.File("test_data/playlist.json"))
	})

	api := NewSpotifyAPI(testBase, client)
	r, err := api.GetPlaylist(id)
	if err != nil {
		t.Error(err)
	}
	if r == nil {
		t.Error("no playlist")
	}

	if r.Id != "3cEYpjA9oz9GiPac4AsH4n" {
		t.Error("invalid playlist id")
	}
}
