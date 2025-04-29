package main

import "net/http"

func (app *application) healthHandler(w http.ResponseWriter, _ *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"version": version,
		"env":     app.config.Env,
	}
	err := app.jsonResponse(w, http.StatusOK, data)
	if err != nil {
		return
	}
}
