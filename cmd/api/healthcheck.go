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
