package model

import "github.com/google/uuid"

type Board struct {
	ID       uuid.UUID    `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	Name     string       `json:"name"`
	Goals    []Goal       `json:"goals"`
	Contents []Content    `json:"contents" gorm:"many2many:board_contents;"`
	Phases   []BoardPhase `json:"phases" `
}
type BoardPhase struct {
	ID      uuid.UUID `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	BoardID uuid.UUID `json:"-" gorm:"type:uuid"`
	PhaseID uuid.UUID `json:"-"`
	Phase   Phase     `json:"phase"`
	Order   uint      `json:"order"`
}
