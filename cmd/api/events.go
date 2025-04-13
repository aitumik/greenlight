package main

import (
	"fmt"
	"greenlight/internal/data"
	"greenlight/internal/validator"
	"log"
	"net/http"
	"time"
)

func (app *application) createEventHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		StartTime   time.Time `json:"start_time"`
		EndTime     time.Time `json:"end_time"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Venue       string    `json:"venue"`
		Tags        []string  `json:"tags"`
		Cover       string    `json:"cover"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	event := &data.Event{
		StartTime:   input.StartTime,
		EndTime:     input.EndTime,
		Title:       input.Title,
		Description: input.Description,
		Venue:       input.Venue,
		Tags:        input.Tags,
		Cover:       input.Cover,
	}

	v := validator.New()
	data.ValidateEvent(v, event)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	created, err := app.models.Events.Insert(event)
	if err != nil {
		log.Fatal(err)
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/events/%d", created.ID))

	data := envelope{
		"event": created,
	}

	err = app.writeJSON(w, http.StatusCreated, data, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showEventHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	event, err := app.models.Events.Get(id)
	if err != nil {

	}

	app.writeJSON(w, http.StatusOK, envelope{"event": event}, nil)
}
