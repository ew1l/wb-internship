package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ew1l/wb-l2/develop/dev11/service"
)

func (h *Handler) result(w http.ResponseWriter, result []*service.Event) {
	w.Header().Set("Content-Type", "application/json")

	data, _ := json.MarshalIndent(struct {
		Result []*service.Event `json:"result"`
	}{
		result,
	}, "", "\t")

	w.Write(data)
}

func (h *Handler) error(w http.ResponseWriter, err string, status int) {
	w.Header().Set("Content-Type", "application/json")

	data, _ := json.MarshalIndent(struct {
		Error string `json:"error"`
	}{
		err,
	}, "", "\t")

	w.WriteHeader(status)
	w.Write(data)
}
