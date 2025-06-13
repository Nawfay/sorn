package models

type Track struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
	AlbumID  int    `json:"album_id"`
	AlbumCover string `json:"album_cover"`
}

type Album struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Artist    string  `json:"artist"`
	ArtistID  int     `json:"artist_id"`
	Cover     string  `json:"cover"`
	CoverBig  string  `json:"cover_big"`
	NbTracks  int     `json:"nb_tracks"`
	Tracks    []Track `json:"tracks"`
}

type Artist struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	NbAlbums      int    `json:"nb_albums"`
	NbFans        int    `json:"nb_fans"`
	PictureBig    string `json:"picture_big"`
	PictureMedium string `json:"picture_medium"`
	PictureSmall  string `json:"picture_small"`
	PictureXL     string `json:"picture_xl"`
}