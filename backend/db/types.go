package db

import "github.com/google/uuid"

type MethodLabel string

const (
	METHOD_LABEL_SINGLE MethodLabel = "METHOD_LABEL_SINGLE"
	METHOD_LABEL_PAIR   MethodLabel = "METHOD_LABEL_PAIR"
	METHOD_LABEL_GROUP  MethodLabel = "METHOD_LABEL_GROUP"
	METHOD_LABEL_PLENUM MethodLabel = "METHOD_LABEL_PLENUM"
)

type Board struct {
	ID       uuid.UUID `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	Name     string    `json:"name"`
	Goals    []Goal    `json:"goals"`
	Contents []Content `json:"contents" gorm:"many2many:board_contents;"`
	Phases   []Phase   `json:"phases" gorm:"many2many:board_phases;"`
}

type Goal struct {
	ID      uuid.UUID `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	BoardID uuid.UUID `json:"boardId"`
	Text    string    `json:"text"`
	Order   uint      `json:"order"`
}

type Phase struct {
	ID      uuid.UUID `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	Title   string    `json:"title"`
	Methods []Method  `gorm:"many2many:phase_methods;" json:"methods"`
}

type Method struct {
	ID          uuid.UUID `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Labels      string    `json:"labels"`
}

type Content struct {
	ID      uuid.UUID `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	BoardID uuid.UUID
	Title   string `json:"title"`
	Text    string `json:"text"`
}

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
	PhaseID  uuid.UUID `json:"phaseId" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	MethodID uuid.UUID `json:"methodId" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	Contents []Content `json:"contents" gorm:"many2many:method_contents;joinForeignKey:phase_id,method_id;"`
	Order    uint      `json:"order"`
}

type MethodContents struct {
	PhaseID   uuid.UUID `json:"phaseId" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	MethodID  uuid.UUID `json:"methodId" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	ContentID uuid.UUID `json:"contentId" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	Order     uint      `json:"order"`
}
