package model

import "github.com/google/uuid"

type Phase struct {
	ID      uuid.UUID `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	Title   string    `json:"title"`
	Methods []Method  `gorm:"many2many:phase_methods;" json:"methods"`
}
