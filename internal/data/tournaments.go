package data

import (
	"database/sql"
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
	Version   int32
}

// TournamentModel defines our tournament Model for DB queries
type TournamentModel struct {
	DB *sql.DB
}

func ValidateTournament(v *validator.Validator, tournament *Tournament) {
	v.Check(tournament.Title != "", "title", "must be provided")
	v.Check(len(tournament.Title) < 500, "title", "must not be 500 bytes long")

	v.Check(tournament.Year != 0, "year", "must be provided")
	v.Check(tournament.Year >= int32(time.Now().Year()), "year", "must not be in the past") // TODO must check that it's not in the past for a tournament

}

// Insert defines our method of inserting into the Db
func (t TournamentModel) Insert(tmt Tournament) error {
	query := `
		INSERT INTO tournaments (title, year)
		VALUES($1, $2)
		RETURNING id, created_at, version
	`
	args := []interface{}{tmt.Title, tmt.Year}

	return t.DB.QueryRow(query, args...).Scan(&tmt.ID, &tmt.CreatedAt, &tmt.Version)
}

// Get defines our method for getting a tournament
func (t TournamentModel) Get(tmt Tournament) (*Tournament, error) {
	return nil, nil
}

// Update defines the method for updating a tournament
func (t TournamentModel) Update(tmt Tournament) error {
	return nil
}

// Delete deines our method for Deleting a tournament from the DB
func (t TournamentModel) Delete(tmt Tournament) error {
	return nil
}
