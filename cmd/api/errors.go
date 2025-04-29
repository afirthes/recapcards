package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Uri: %s Internal server error: %s", r.URL.Path, err)
	writeJSONError(w, http.StatusInternalServerError, "The server encountered an internal error")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Uri: %s Bad request: %s", r.URL.Path, err)
	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Uri: %s Not found: %s", r.URL.Path, err)
	writeJSONError(w, http.StatusNotFound, "Not found")
}
