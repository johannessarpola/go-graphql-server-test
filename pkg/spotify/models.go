package spotify

type Playlist struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Id          string `json:"id"`
}

type GetFeaturedPlaylists struct {
	//Playlists Playlists `json:"playlists"`

	Playlists struct {
		Limit  int        `json:"limit"`
		Offset int        `json:"offset"`
		Total  int        `json:"total"`
		Items  []Playlist `json:"items"`
	} `json:"playlists"`
}

type Track struct {
	Uri        string `json:"uri"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	DurationMs int    `json:"duration_ms"`
	Explicit   bool   `json:"explicit"`
}

type GetTracks struct {
	Items []struct {
		Track Track `json:"track"`
	} `json:"items"`
}

type AddTracksPayload struct {
	Id   string   `json:"id"`
	Uris []string `json:"uris"`
}

type AddTracksRequest struct {
	Position *int     `json:"position,omitempty"`
	Uris     []string `json:"uris"`
}

type SnapshotOrError struct {
	SnapshotId string `json:"snapshot_id"`
	Error      string `json:"error"`
}
