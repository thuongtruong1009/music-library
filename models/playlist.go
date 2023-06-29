package models

type Playlist struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Tracks []string `json:"tracks"`
	Duration float32 `json:"duration"`
}