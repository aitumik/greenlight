package main

import (
	"fmt"
	"net/http"
)

func (app *application) createMovieHandler(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintln(w,"create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"show a movie")
}

