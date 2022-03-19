package main

import (
	"errors"
	"fmt"
	"greenlight/internal/data"
	"greenlight/internal/validator"
	"net/http"
)

// BODY='{"title" : "Moana","year" : 2016,"runtime" : 107,"genres" : ["animation","adventure"]}'

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	// Decode the request
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}
	// initialize a new validator instance
	v := validator.New()
	data.ValidateMovie(v, movie)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Movies.Insert(movie)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/movies/%d", movie.ID))

	data := envelope{
		"movie": movie,
	}
	err = app.writeJSON(w, http.StatusCreated, data, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	movie, err := app.models.Movies.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		// display json formatted error to the client
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
	}

	movie, err := app.models.Movies.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// We only expect user to change these fields
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	err = app.readJSON(w,r,&input)
	if err != nil {
		app.badRequestResponse(w,r,err)
		return
	}

	// Copy the values from the input to the movie
	movie.Title = input.Title
	movie.Year  = input.Year
	movie.Runtime = input.Runtime
	movie.Genres = input.Genres

	v := validator.New()

	if data.ValidateMovie(v,movie); !v.Valid() {
		app.failedValidationResponse(w,r,v.Errors)
		return
	}

	err = app.models.Movies.Update(movie)
	if err != nil {
		app.serverErrorResponse(w,r,err)
		return
	}

	data := envelope{
		"movie" : movie,
	}
	err = app.writeJSON(w,http.StatusCreated,data,nil)
	if err != nil {
		app.serverErrorResponse(w,r,err)
	}
}

func (app *application) deleteMovieHandler(w http.ResponseWriter,r *http.Request) {
	id,err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w,r)
		return
	}
	err = app.models.Movies.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err,data.ErrRecordNotFound):
			app.notFoundResponse(w,r)
		default:
			app.serverErrorResponse(w,r,err)
		}
		return
	}

	err = app.writeJSON(w,http.StatusOK,envelope{"message" : "movie deleted successfully"},nil)
	if err != nil {
		app.serverErrorResponse(w,r,err)
	}
}
