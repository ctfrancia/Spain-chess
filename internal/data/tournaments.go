package data

import (
	"time"

	"jaque.ctfrancia.com/internal/validator"
)

// Tournament describes the shape of a tournament
type Tournament struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Year      int32     `json:"year"`
	Month     int8      `json:"month"`
	Day       int32     `json:"day"`
}

func ValidateTournament(v *validator.Validator, tournament *Tournament) {
	v.Check(tournament.Title != "", "title", "must be provided")
	v.Check(len(tournament.Title) < 500, "title", "must not be 500 bytes long")

	v.Check(tournament.Year != 0, "year", "must be provided")
	v.Check(tournament.Year >= int32(time.Now().Year()), "year", "must not be in the past") // TODO must check that it's not in the past for a tournament

}
