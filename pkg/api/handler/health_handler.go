package handler

import (
	"github.com/go-chi/render"
	"net/http"
)

type healthAPIHandler struct{}

type healthResponse struct {
	Status string `json:"status"`
}

func NewAPIHealthHandler() *healthAPIHandler {
	handler := &healthAPIHandler{}
	return handler
}

func (h *healthAPIHandler) IndexController(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, healthResponse{Status: "ok"})
}
