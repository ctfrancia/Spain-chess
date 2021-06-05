package main

import (
	"fmt"
	"net/http"
	"time"

	"jaque.ctfrancia.com/internal/data"
	"jaque.ctfrancia.com/internal/validator"
)

func (app *application) createTournamentHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
		Year  int32  `json:"year"`
		// People []string `json:"people"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	tournament := &data.Tournament{
		Title: input.Title,
		Year:  input.Year,
	}

	v := validator.New()

	if data.ValidateTournament(v, tournament); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showTournamentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	tournament := data.Tournament{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "tournament 123",
		Year:      2021,
		Month:     5,
		Day:       27,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"tournament": tournament}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
