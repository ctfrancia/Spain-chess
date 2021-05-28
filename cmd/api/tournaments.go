package main

import (
	"fmt"
	"net/http"
	"time"

	"jaque.ctfrancia.com/internal/data"
)

func (app *application) createTournamentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new tournament")
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
