package spotify

type Tracks struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type Playlist struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	SnapshotId  string `json:"snapshot_id"`
	Tracks      Tracks `json:"tracks"`
	Uri         string `json:"uri"`
}

type Playlists struct {
	Items  []Playlist `json:"items"`
	Limit  int        `json:"limit"`
	Offset int        `json:"offset"`
	Total  int        `json:"total"`
}

type PlaylistResponse struct {
	Playlists Playlists `json:"playlists"`
}
