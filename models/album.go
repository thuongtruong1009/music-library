package models

type Album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Duration float32 `json:"duration"`
	Year int `json:"year"`

	ArtistID string `json:"artist_id"`
}

