package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Retreive the "id" URL param form the currrent request context, then convert it to
// an integer and return it. If the operation isn"t succesful, return 0 and an error
func (app *Application) readIDParams(r *http.Request) (int64, error) {
	// Get the params
	params := httprouter.ParamsFromContext(r.Context())

	// Convert id to int
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		// returning the error
		return 0, errors.New("invalid id parameter")
	}

	// return the valid value
	return id, nil
}

type envelope map[string]interface{}

// Takes a data interface{} and writes it inot the response writer
// with the given headers and the status code
func (app *Application) writeJSON(data envelope, w http.ResponseWriter, httpStatus int, headers http.Header)  error {
  // convert data to json 
  json, err := json.Marshal(data)
  if err != nil {
    return err
  }

  // write the json data 
  w.Write(json)

  // put all the headers into the response 
  for k, v := range headers{
    w.Header()[k] = v
  }

  // write the headers 
  w.Header().Add("Content-type", "application/json")
  w.WriteHeader(httpStatus)

  // return nil as everything is completed without error
  return nil
}
