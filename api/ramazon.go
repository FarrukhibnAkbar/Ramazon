package api

import (
	"net/http"
	"ramazon/models"
)

func (api *api) SetPrayTime(w http.ResponseWriter, r *http.Request) {
	var body models.Ramazon
	if err := BodyParser(r, &body); err != nil {
		HandleErrorResponse(w, 500, "body parse error")
		return
	}

	err := api.ramazonService.SetPrayTime(r.Context(), body)
	if err != nil {
		HandleBadRequestErrWithMessage(w, err, "error set pray time")
		return
	}

	WriteJSONWithSuccess(w, 201)
}
