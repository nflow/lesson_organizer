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
	Rank    uint      `json:"rank"`
}

type BoardPhase struct {
	ID      uuid.UUID     `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	BoardID uuid.UUID     `json:"-" gorm:"type:uuid"`
	PhaseID uuid.UUID     `json:"-" gorm:"type:uuid"`
	Phase   Phase         `json:"phase"`
	Methods []BoardMethod `json:"methods"`
	Rank    uint          `json:"rank"`
}

type BoardMethod struct {
	ID           uuid.UUID      `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	BoardPhaseID uuid.UUID      `json:"-" gorm:"type:uuid"`
	MethodID     uuid.UUID      `json:"-" gorm:"type:uuid"`
	Method       Method         `json:"method"`
	Contents     []BoardContent `json:"contents"`
	Rank         uint           `json:"rank"`
}

// TODO: Create seperate DTOs for everything and move to another package
type CreateContentDto struct {
	Text string `json:"text"`
}

type BoardContent struct {
	ID            uuid.UUID `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key;"`
	BoardID       uuid.UUID `json:"boardId" gorm:"type:uuid"`
	BoardMethodID uuid.UUID `json:"methodId" gorm:"type:uuid"`
	Text          string    `json:"text"`
	Rank          uint      `json:"rank"`
}

type MovePhaseDto struct {
	PhaseID      uuid.UUID `json:"phaseId"`
	AfterPhaseID uuid.UUID `json:"afterPhaseId"`
}

type MoveMethodDto struct {
	MethodID      uuid.UUID `json:"methodId"`
	ParentPhaseID uuid.UUID `json:"parentPhaseId"`
	AfterMethodID uuid.UUID `json:"afterMethodId"`
}

type MoveContentDto struct {
	ContentID      uuid.UUID `json:"contentId"`
	ParentMethodID uuid.UUID `json:"parentMethodId"`
	ParentBoardID  uuid.UUID `json:"parentBoardId"`
	AfterContentID uuid.UUID `json:"afterContentId"`
}
