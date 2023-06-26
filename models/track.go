package models

type Track struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Year int `json:"year"`
	Duration float32 `json:"duration"`
	
	GenreID string `json:"genre_id"`
	ArtistID string `json:"artist_id"`
}