package main

// Routes for the greenlight application
import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// Initialize the router
	router := httprouter.New()

	// NotFound error handler
	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	// MethodNotAllowed error handler
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet,"/v1/movies",app.listMoviesHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)
	router.HandlerFunc(http.MethodPatch,"/v1/movies/:id",app.updateMovieHandler)
	router.HandlerFunc(http.MethodDelete,"/v1/movies/:id",app.deleteMovieHandler)

	return router
}
