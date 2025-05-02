package main

import "net/http"

// healthcheckHandler godoc
//
//	@Summary		Healthcheck
//	@Description	Healthcheck endpoint
//	@Tags			ops
//	@Produce		json
//	@Success		200	{object}	string	"ok"
//	@Router			/health [get]
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
