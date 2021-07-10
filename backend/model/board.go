package model

import "github.com/google/uuid"

type Board struct {
	ID       uuid.UUID `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	Name     string    `json:"name"`
	Goals    []Goal    `json:"goals"`
	Contents []Content `json:"contents" gorm:"many2many:board_contents;"`
	Phases   []Phase   `json:"phases" gorm:"many2many:board_phases;"`
}
