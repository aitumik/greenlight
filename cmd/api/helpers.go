package main

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (app *application) readIDParam(r *http.Request) (int64,error) {
	params := httprouter.ParamsFromContext(r.Context())

	id,err := strconv.ParseInt(params.ByName("id"),10,64)
	if err != nil || id < 1 {
		return 0,errors.New("invalid parameter")
	}
	return id,nil
}

func (app *application) writeJSON(w http.ResponseWriter,status int,data interface{},headers http.Header) error {
	js,err := json.MarshalIndent(data,"","\t")
	if err != nil {
		return err
	}

	js = append(js,'\n')

	for key,value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}
