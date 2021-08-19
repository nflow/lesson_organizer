package model

import "github.com/google/uuid"

type Board struct {
	ID       uuid.UUID      `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	Name     string         `json:"name"`
	Goals    []BoardGoal    `json:"goals"`
	Contents []BoardContent `json:"contents"`
	Phases   []BoardPhase   `json:"phases"`
}

type BoardGoal struct {
	ID      uuid.UUID `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	BoardID uuid.UUID `json:"-" gorm:"type:uuid"`
	Text    string    `json:"text"`
	Order   uint      `json:"order"`
}

type BoardPhase struct {
	ID      uuid.UUID     `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	BoardID uuid.UUID     `json:"-" gorm:"type:uuid"`
	PhaseID uuid.UUID     `json:"-" gorm:"type:uuid"`
	Phase   Phase         `json:"phase"`
	Methods []BoardMethod `json:"methods"`
	Order   uint          `json:"order"`
}

type BoardMethod struct {
	ID           uuid.UUID      `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	BoardPhaseID uuid.UUID      `json:"-" gorm:"type:uuid"`
	MethodID     uuid.UUID      `json:"-" gorm:"type:uuid"`
	Method       Method         `json:"method"`
	Contents     []BoardContent `json:"contents"`
	Order        uint           `json:"order"`
}

type BoardContent struct {
	ID            uuid.UUID `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	BoardID       uuid.UUID `json:"-" gorm:"type:uuid"`
	BoardMethodID uuid.UUID `json:"-" gorm:"type:uuid"`
	Text          string    `json:"text"`
	Order         uint      `json:"order"`
}
