package models

type Playlist struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Duration float32 `json:"duration"`
}