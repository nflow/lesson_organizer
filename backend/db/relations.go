package db

import (
	"github.com/google/uuid"
	"github.com/nflow/lesson_organizer/model"
)

type BoardContents struct {
	BoardID   uuid.UUID `json:"boardId" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	ContentID uuid.UUID `json:"contentId" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	Order     uint      `json:"order"`
}

type BoardPhases struct {
	BoardID uuid.UUID `json:"boardId" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	PhaseID uuid.UUID `json:"phaseId" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	Order   uint      `json:"order"`
}

type PhaseMethods struct {
	PhaseID  uuid.UUID       `json:"phaseId" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	MethodID uuid.UUID       `json:"methodId" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	Contents []model.Content `json:"contents" gorm:"many2many:method_contents;joinForeignKey:phase_id,method_id;"`
	Order    uint            `json:"order"`
}

type MethodContents struct {
	PhaseID   uuid.UUID `json:"phaseId" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	MethodID  uuid.UUID `json:"methodId" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	ContentID uuid.UUID `json:"contentId" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	Order     uint      `json:"order"`
}
