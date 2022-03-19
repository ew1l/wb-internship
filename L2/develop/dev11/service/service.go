package service

import (
	"errors"
	"sync"
	"time"
)

// Service ...
type Service interface {
	CreateEvent(userID, ID, date, description string) (*Event, error)
	UpdateEvent(userID, ID, date, description string) (*Event, error)
	DeleteEvent(userID, ID string) (*Event, error)
	DateEvents(userID, date string, duration time.Duration) ([]*Event, error)

	userExists(userID string) bool
	eventExists(userID, ID string) bool
}

var (
	errUserIDNotFound       = errors.New("user ID not found")
	errEventIDNotFound      = errors.New("event ID not found")
	errEventIDAlreadyExists = errors.New("event ID already exists")
)

type service struct {
	events map[string][]*Event
	mu     *sync.Mutex
}

// NewService service's structure constructor
func NewService() Service {
	return &service{
		events: make(map[string][]*Event, 0),
		mu:     new(sync.Mutex),
	}
}

// CreateEvent to create an event
func (s *service) CreateEvent(userID, ID, date, description string) (*Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.eventExists(userID, ID) {
		return nil, errEventIDAlreadyExists
	}

	event := &Event{
		ID:          ID,
		Date:        date,
		Description: description,
	}

	s.events[userID] = append(s.events[userID], event)

	return event, nil
}

// UpdateEvent to update an event
func (s *service) UpdateEvent(userID, ID, date, description string) (*Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.userExists(userID) {
		return nil, errUserIDNotFound
	}

	if !s.eventExists(userID, ID) {
		return nil, errEventIDNotFound
	}

	event := new(Event)
	for i, e := range s.events[userID] {
		if e.ID == ID {
			s.events[userID][i].Description = description
			s.events[userID][i].Date = date
			event = e
			break
		}
	}

	return event, nil
}

// DeleteEvent to delete an event
func (s *service) DeleteEvent(userID, ID string) (*Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.userExists(userID) {
		return nil, errUserIDNotFound
	}

	if !s.eventExists(userID, ID) {
		return nil, errEventIDNotFound
	}

	event := new(Event)
	for i, e := range s.events[userID] {
		if e.ID == ID {
			event = e
			s.events[userID] = append(s.events[userID][:i], s.events[userID][i+1:]...)
			break
		}
	}

	return event, nil
}

// DateEvents to get events for a date
func (s *service) DateEvents(userID, date string, duration time.Duration) ([]*Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.userExists(userID) {
		return nil, errUserIDNotFound
	}

	start, _ := time.Parse("2006-01-02", date)
	end := start.Add(duration)

	events := make([]*Event, 0)
	for _, event := range s.events[userID] {
		parsed, _ := time.Parse("2006-01-02", event.Date)

		if (parsed.After(start) || parsed == start) && parsed.Before(end) {
			events = append(events, event)
		}
	}

	return events, nil
}

func (s *service) userExists(userID string) bool {
	if _, ok := s.events[userID]; !ok {
		return false
	}

	return true
}

func (s *service) eventExists(userID, ID string) bool {
	for _, event := range s.events[userID] {
		if event.ID == ID {
			return true
		}
	}

	return false
}
