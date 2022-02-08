package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (app *application) createMovieHandler(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintln(w,"create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id,err := strconv.ParseInt(params.ByName("id"),10,64)
	if err != nil || id < 1 {
		http.NotFound(w,r)
		return
	}
	fmt.Fprintf(w,"showing the details of movie %d\n",id)
}

