package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	healthCkeckData := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	// get the json representation of the healthcheck data in bytes slice
	err := app.writeJson(w, http.StatusOK, healthCkeckData, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error : The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}
}

func (app *application) writeJson(w http.ResponseWriter, status int, data any, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	for key, value := range headers {
		w.Header()[key] = value
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *application) exampleHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"hello": "world",
	}
	// Set the "Content-Type: application/json" header on the response.
	w.Header().Set("Content-Type", "application/json")
	// Use the json.NewEncoder() function to initialize a json.Encoder instance that
	// writes to the http.ResponseWriter. Then we call its Encode() method, passing in
	// the data that we want to encode to JSON (which in this case is the map above). If
	// the data can be successfully encoded to JSON, it will then be written to our
	// http.ResponseWriter.
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
