package main

import (
	"fmt"
	"greenlight/internal/data"
	"greenlight/internal/validator"
	"net/http"
	"time"
)

// BODY='{"title" : "Moana","year" : 2016,"runtime" : 107,"genres" : ["animation","adventure"]}'

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime data.Runtime    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	// Decode the request
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	movie := &data.Movie{
		Title: input.Title,
		Year: input.Year,
		Runtime: input.Runtime,
		Genres: input.Genres,
	}
	// initialize a new validator instance
	v := validator.New()
	data.ValidateMovie(v,movie)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	// Dump the contents of the input struct in an HTTP response
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		// display json formatted error to the client
		app.serverErrorResponse(w, r, err)
	}
	return
}
