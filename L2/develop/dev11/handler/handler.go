package handler

import (
	"net/http"
	"time"

	"github.com/ew1l/wb-l2/develop/dev11/service"
)

// Handler for service
type Handler struct {
	service.Service
}

// NewHandler handler's structure constructor
func NewHandler(s service.Service) *Handler {
	return &Handler{
		Service: s,
	}
}

// CreateEvent handles the request to create an event
func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	eventDTO, err := h.parse(r)
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
	}

	if err := h.validate(eventDTO); err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	created, err := h.Service.CreateEvent(eventDTO.UserID, eventDTO.ID, eventDTO.Date, eventDTO.Description)
	if err != nil {
		h.error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	h.result(w, []*service.Event{created})
}

// UpdateEvent handles the request to update an event
func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	eventDTO, err := h.parse(r)
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
	}

	if err := h.validate(eventDTO); err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := h.Service.UpdateEvent(eventDTO.UserID, eventDTO.ID, eventDTO.Date, eventDTO.Description)
	if err != nil {
		h.error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	h.result(w, []*service.Event{updated})
}

// DeleteEvent handles the request to delete an event
func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	eventDTO, err := h.parse(r)
	if err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
	}

	deleted, err := h.Service.DeleteEvent(eventDTO.UserID, eventDTO.ID)
	if err != nil {
		h.error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	h.result(w, []*service.Event{deleted})
}

// EventsForDay handles the request to get events for day
func (h *Handler) EventsForDay(w http.ResponseWriter, r *http.Request) {
	eventDTO, _ := h.parse(r)
	if err := h.validate(eventDTO); err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	events, err := h.Service.DateEvents(eventDTO.UserID, eventDTO.Date, 24*time.Hour)
	if err != nil {
		h.error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	h.result(w, events)
}

// EventsForWeek handles the request to get events for week
func (h *Handler) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	eventDTO, _ := h.parse(r)
	if err := h.validate(eventDTO); err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}
	events, err := h.Service.DateEvents(eventDTO.UserID, eventDTO.Date, 7*24*time.Hour)
	if err != nil {
		h.error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	h.result(w, events)
}

// EventsForMonth handles the request to get events for month
func (h *Handler) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	eventDTO, _ := h.parse(r)
	if err := h.validate(eventDTO); err != nil {
		h.error(w, err.Error(), http.StatusBadRequest)
		return
	}

	events, err := h.Service.DateEvents(eventDTO.UserID, eventDTO.Date, 4*7*24*time.Hour)
	if err != nil {
		h.error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	h.result(w, events)
}
