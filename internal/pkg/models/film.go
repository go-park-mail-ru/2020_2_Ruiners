package models

type Film struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `'json:"description"'`
	MainGenre string `'json:"main_genre"'`
	YoutubeLink	string `'json:"youtube_link"'`
	BigImg string `'json:"big_img"'`
	SmallImg string `'json:"small_img"'`
	Year int `json:"year"`
	Country string `json:"country"`
}
