package model

import "github.com/google/uuid"

type Goal struct {
	ID      uuid.UUID `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	BoardID uuid.UUID `json:"boardId"`
	Text    string    `json:"text"`
	Order   uint      `json:"order"`
}
