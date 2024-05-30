package spotify

type Playlist struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	SnapshotId  string `json:"snapshot_id"`
	Uri         string `json:"uri"`
	Id          string `json:"id"`
}

type GetFeaturedPlaylists struct {
	//Playlists Playlists `json:"playlists"`

	Playlists struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Total  int `json:"total"`
		Items  []struct {
			Id          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"items"`
	} `json:"playlists"`
}

type Track struct {
	Uri        string `json:"uri"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	DurationMs int    `json:"duration_ms"`
	Explicit   bool   `json:"explicit"`
}

type GetPlaylist struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tracks      []struct {
		Items []struct {
			Track Track `json:"track"`
		} `json:"items"`
	} `json:"tracks"`
}

type GetTracks struct {
	Items []struct {
		Track Track `json:"track"`
	} `json:"items"`
}
