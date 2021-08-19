package model

import "github.com/google/uuid"

type Phase struct {
	ID    uuid.UUID `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	Title string    `json:"title"`
}

type PhaseIdentifier struct {
	ID uuid.UUID `json:"id" validate:"required,uuid"`
}
