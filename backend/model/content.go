package model

import "github.com/google/uuid"

type Content struct {
	ID      uuid.UUID `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	BoardID uuid.UUID
	Title   string `json:"title"`
	Text    string `json:"text"`
}
