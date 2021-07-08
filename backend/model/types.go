package model

import "github.com/google/uuid"

type MethodLabel string

const (
	METHOD_LABEL_SINGLE = "METHOD_LABEL_SINGLE"
	METHOD_LABEL_PAIR   = "METHOD_LABEL_PAIR"
	METHOD_LABEL_GROUP  = "METHOD_LABEL_GROUP"
	METHOD_LABEL_PLENUM = "METHOD_LABEL_PLENUM"
)

type Board struct {
	ID       uuid.UUID `json:ID validate:"required,uuid" grom:"primaryKey"`
	Name     string    `json:name`
	Goals    []Goal    `json:goals`
	Contents []Content `json:contents`
	Phases   []Phase   `json:phases gorm:"many2many:board_phases;"`
}

type Goal struct {
	ID    uuid.UUID `json:ID validate:"required,uuid" grom:"primaryKey"`
	Text  string    `json:text`
	Order uint      `json:order`
}

type Phase struct {
	ID      uuid.UUID `json:ID validate:"required,uuid" grom:"primaryKey"`
	Title   string    `json:title`
	Methods []Method  `json:methods gorm:"many2many:phase_methods;"`
}

type Method struct {
	ID          uuid.UUID     `json:ID validate:"required,uuid" grom:"primaryKey"`
	Title       string        `json:title`
	Description string        `json:description`
	Labels      []MethodLabel `json:labels`
	Contents    []Content     `json:contents`
}

type Content struct {
	ID    uuid.UUID `json:ID validate:"required,uuid" grom:"primaryKey"`
	Title string    `json:title`
	Text  string    `json:text`
	Order uint      `json:order`
}
