package model

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type MethodLabel string

const (
	METHOD_LABEL_SINGLE MethodLabel = "METHOD_LABEL_SINGLE"
	METHOD_LABEL_PAIR   MethodLabel = "METHOD_LABEL_PAIR"
	METHOD_LABEL_GROUP  MethodLabel = "METHOD_LABEL_GROUP"
	METHOD_LABEL_PLENUM MethodLabel = "METHOD_LABEL_PLENUM"
)

type Method struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Category    string         `json:"category"`
	Labels      pq.StringArray `json:"labels" gorm:"type:text[]"`
}

type MethodIdentifier struct {
	ID uuid.UUID `json:"id" validate:"required,uuid"`
}
