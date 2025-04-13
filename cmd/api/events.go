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
		// TODO : we need to know the error type so as to show correct feedback to user
		app.notFoundResponse(w, r)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"event": event}, nil)
}

func (app *application) listEventsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string
		Description string
		Venue       string
		Location    string
		Tags        []string
		data.Filters
	}

	q := r.URL.Query()

	input.Title = app.readString(q, "title", "")
	input.Description = app.readString(q, "description", "")
	input.Venue = app.readString(q, "venue", "")
	input.Location = app.readString(q, "location", "Nairobi")
	input.Tags = app.readCSV(q, "tags", []string{})

	v := validator.New()

	input.Filters.Page = app.readInt(q, "page", 1, v)
	input.Filters.PageSize = app.readInt(q, "page_size", 10, v)
	input.Filters.Sort = app.readString(q, "sort", "id")

	input.Filters.SortSafelist = []string{"title", "description", "-title", "-description"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	events, metadata, err := app.models.Events.GetAll(input.Location, input.Title, input.Description, input.Tags, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"events": events, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
