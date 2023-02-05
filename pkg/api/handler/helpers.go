package handler

import (
	"fmt"
	"github.com/go-chi/render"
	"net/http"
)

type errorPayload struct {
	Message string `json:"error"`
}

func responseCustomRender(statusCode int, msg interface{}, w http.ResponseWriter, r *http.Request) {
	render.Status(r, statusCode)
	render.JSON(w, r, msg)
}

func responseOKRender(msg interface{}, w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	if msg == nil {
		render.JSON(w, r, map[string]string{"status": "OK"})
	} else {
		render.JSON(w, r, msg)
	}
}

func responseBadRequestRender(msg string, w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusBadRequest)
	render.JSON(w, r, &errorPayload{Message: msg})
}

func responseUnprocessableEntityRender(msg string, w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusUnprocessableEntity)
	render.JSON(w, r, &errorPayload{Message: msg})
}

func responseInternalErrorRender(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusInternalServerError)
	render.JSON(w, r, &errorPayload{Message: "Internal server error"})
}

func responseForbiddenRender(msg string, w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusForbidden)
	render.JSON(w, r, &errorPayload{Message: msg})
}

func responseNotFound(entityId string, w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusNotFound)
	render.JSON(w, r, &errorPayload{Message: fmt.Sprintf("entity %v not found", entityId)})
}

func responseNoContent(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusNoContent)
	render.JSON(w, r, &errorPayload{Message: "No content"})
}
