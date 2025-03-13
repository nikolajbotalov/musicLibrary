package domain

import "time"

type Song struct {
	ID          string    `json:"id"`
	ArtistID    string    `json:"artists_id"`
	Title       string    `json:"title"`
	ReleaseDate *string   `json:"release_date"`
	Text        *string   `json:"text"`
	Link        *string   `json:"link"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
