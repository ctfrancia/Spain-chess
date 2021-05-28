package data

import "time"

// Tournament describes the shape of a tournament
type Tournament struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Year      int32     `json:"year"`
	Month     int8      `json:"month"`
	Day       int32     `json:"day"`
}
