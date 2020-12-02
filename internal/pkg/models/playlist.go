package models

type Playlist struct {
	Id     int        `json:"id"`
	Title  string     `json:"title"`
	Films  *FilmCards `json:"films"`
	UserId int        `json:"user_id"`
}

//easyjson:json
type Playlists []Playlist
