package service

// Event structure
type Event struct {
	ID          string `json:"id"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

// EventDTO structure for passing data to the service
type EventDTO struct {
	UserID      string `json:"user_id"`
	ID          string `json:"id"`
	Date        string `json:"date"`
	Description string `json:"description"`
}
