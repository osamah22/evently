package models

import "github.com/google/uuid"

type Event struct {
	ID    uuid.UUID
	Title string
}
