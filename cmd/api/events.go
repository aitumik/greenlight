package main

import (
	"fmt"
	"greenlight/internal/data"
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

	// since its a pointer I expect the event to contain id
	app.models.Events.Insert(event)

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/events/%d", event.ID))

	data := envelope{
		"event": event,
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
