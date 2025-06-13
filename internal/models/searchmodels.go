package models


// APIResponse represents the structure of the API response
type AlbumResponse struct {
	Data  []AlbumRaw `json:"data"`
	Total int     `json:"total"`
	Next  string  `json:"next"`
}

// Album represents the structure of an album in the API response
type AlbumRaw struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Link          string  `json:"link"`
	Cover         string  `json:"cover"`
	CoverSmall    string  `json:"cover_small"`
	CoverMedium   string  `json:"cover_medium"`
	CoverBig      string  `json:"cover_big"`
	CoverXL       string  `json:"cover_xl"`
	MD5Image      string  `json:"md5_image"`
	GenreID       int     `json:"genre_id"`
	NumberOfTracks int    `json:"nb_tracks"`
	RecordType    string  `json:"record_type"`
	Tracklist     string  `json:"tracklist"`
	ExplicitLyrics bool   `json:"explicit_lyrics"`
	Artist        ArtistRaw  `json:"artist"`
	Type          string  `json:"type"`
}




type TrackResponse struct {
	Data  []TrackRaw `json:"data"`
	Total int     `json:"total"`
}

// Track represents the structure of a track in the API response
type TrackRaw struct {
	ID                   int     `json:"id"`
	Readable             bool    `json:"readable"`
	Title                string  `json:"title"`
	TitleShort           string  `json:"title_short"`
	TitleVersion         string  `json:"title_version"`
	Link                 string  `json:"link"`
	Duration             int     `json:"duration"`
	Rank                 int     `json:"rank"`
	ExplicitLyrics       bool    `json:"explicit_lyrics"`
	ExplicitContentLyrics int    `json:"explicit_content_lyrics"`
	ExplicitContentCover int     `json:"explicit_content_cover"`
	Preview              string  `json:"preview"`
	MD5Image             string  `json:"md5_image"`
	Artist               ArtistRaw  `json:"artist"`
	Album                TrackAlbumRaw   `json:"album"`
	Type                 string  `json:"type"`
}


// Album represents the structure of an album in the API response
type TrackAlbumRaw struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Cover         string `json:"cover"`
	CoverSmall    string `json:"cover_small"`
	CoverMedium   string `json:"cover_medium"`
	CoverBig      string `json:"cover_big"`
	CoverXL       string `json:"cover_xl"`
	MD5Image      string `json:"md5_image"`
	Tracklist     string `json:"tracklist"`
	Type          string `json:"type"`
}

// Artist represents the structure of an artist in the API response
type ArtistRaw struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Link           string `json:"link"`
	Picture        string `json:"picture"`
	PictureSmall   string `json:"picture_small"`
	PictureMedium  string `json:"picture_medium"`
	PictureBig     string `json:"picture_big"`
	PictureXL      string `json:"picture_xl"`
	Tracklist      string `json:"tracklist"`
	NbAlbums       int    `json:"nb_albums"`
	Type           string `json:"type"`
}


// TrackResponse represents the structure of the API response
type AlbumTrackResponse struct {
	Data  []TrackRaw `json:"data"`
	Total int        `json:"total"`
}

// TrackRaw represents the structure of a track in the API response
type AlbumTrackRaw struct {
	ID                   int        `json:"id"`
	Readable             bool       `json:"readable"`
	Title                string     `json:"title"`
	TitleShort           string     `json:"title_short"`
	TitleVersion         string     `json:"title_version"`
	ISRC                 string     `json:"isrc"`
	Link                 string     `json:"link"`
	Duration             int        `json:"duration"`
	TrackPosition        int        `json:"track_position"`
	DiskNumber           int        `json:"disk_number"`
	Rank                 int        `json:"rank"`
	ExplicitLyrics       bool       `json:"explicit_lyrics"`
	ExplicitContentLyrics int       `json:"explicit_content_lyrics"`
	ExplicitContentCover int        `json:"explicit_content_cover"`
	Preview              string     `json:"preview"`
	MD5Image             string     `json:"md5_image"`
	Artist               ArtistRaw  `json:"artist"`
	Type                 string     `json:"type"`
}

// ArtistRaw represents the structure of an artist in the API response
type AlbumArtistRaw struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Tracklist string `json:"tracklist"`
	Type      string `json:"type"`
}



