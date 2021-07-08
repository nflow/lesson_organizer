package model

import "github.com/google/uuid"

type BoardPhases struct {
	BoardID uuid.UUID `json:boardId validate:"required,uuid" grom:"primaryKey"`
	PhaseID uuid.UUID `json:phaseId validate:"required,uuid" grom:"primaryKey"`
	Order   uint      `json:order`
}

type PhaseMethods struct {
	PhaseID  uuid.UUID `json:phaseId validate:"required,uuid" grom:"primaryKey"`
	MethodID uuid.UUID `json:methodId validate:"required,uuid" grom:"primaryKey"`
	Order    uint      `json:order`
}
