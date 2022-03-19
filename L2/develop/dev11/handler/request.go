package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/ew1l/wb-l2/develop/dev11/service"
)

func (h *Handler) parse(r *http.Request) (*service.EventDTO, error) {
	eventDTO := new(service.EventDTO)

	switch r.Method {
	case "POST":
		err := json.NewDecoder(r.Body).Decode(eventDTO)
		if err != nil {
			return nil, err
		}
	case "GET":
		eventDTO.UserID = r.URL.Query().Get("user_id")
		eventDTO.Date = r.URL.Query().Get("date")
	}

	return eventDTO, nil
}

func (h *Handler) validate(dto *service.EventDTO) error {
	if _, err := time.Parse("2006-01-02", dto.Date); err != nil {
		return errors.New("incorrect date")
	}

	return nil
}
