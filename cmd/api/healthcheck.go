package main

import (
	"encoding/json"
	"net/http"
)

// Define an envelope type.
type envelope map[string]any

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Declare an envelope map containing the data for the response. Notice that the way we've constructed this means the environment
	// and version data will now be nested under a system_info key in the JSON response.
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}
	err := app.writeJson(w, http.StatusOK, env, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

// Change the data parameter to have the type envelope instead of any.
func (app *application) writeJson(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	js = append(js, '\n')
	for key, value := range headers {
		w.Header()[key] = value
	}
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
