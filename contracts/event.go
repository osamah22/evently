package contracts

import (
	"time"

	"github.com/google/uuid"
)

type EventResponse struct {
	ID                uuid.UUID `json:"id"`
	Title             string    `json:"title"`
	Description       string    `json:"desrciption"`
	OrganizerID       uuid.UUID `json:"organizer_id"`
	Category          string    `json:"category"`
	Location          string    `json:"location"`
	ImageURL          *string   `json:"image_url,omitempty"`
	MaxCapacity       int       `json:"max_capacity"`
	Status            string    `json:"status"`
	StartingTime      time.Time `json:"starting_time"`
	DurationInMinutes int64     `json:"duration_in_minutes"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type CreateEventRequest struct {
}
