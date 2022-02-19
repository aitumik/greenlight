package main

import (
	"fmt"
	"net/http"
)

func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

// errorResponse is a generic helper method for sending
// json encoded responses to the client
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	// Create an envelope to envelope the error message
	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(w, err)

	message := "the server encountered an error and couldn't process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// The notFoundResponse() method will be used to send a 404 Not Found status code and
// JSON response to the client
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource couldn't be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// The methodNotAllowedResponse() method will be used to send a 405 Method Not Allowed
// status code and JSON response to the client.
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s is not supported by the resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
