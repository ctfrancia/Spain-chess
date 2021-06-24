package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Tournaments TournamentModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Tournaments: TournamentModel{DB: db},
	}
}
