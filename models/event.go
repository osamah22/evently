package models

import (
	"time"

	"github.com/google/uuid"
)

type EventStatus string

const (
	StatusDraft     EventStatus = "DRAFT"
	StatusPublished EventStatus = "PUBLISHED"
	StatusCancelled EventStatus = "CANCELLED"
)

type Event struct {
	ID                uuid.UUID
	Title             string
	Description       string
	OrganizerID       uuid.UUID
	CategoryID        int32 `json:"category_id"`
	Location          string
	ImageURL          *string // pointer since nullable
	MaxCapacity       *int    // pointer since nullable
	Status            EventStatus
	StartingTime      time.Time
	DurationInMinutes int64
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
