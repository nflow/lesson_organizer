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
	Id   uuid.UUID `json:id validate:"required,uuid"`
	Name string
}

type Goal struct {
	Id   uuid.UUID `json:id validate:"required,uuid"`
	Text string    `json:text`
}

type Phase struct {
	Id    uuid.UUID `json:id validate:"required,uuid"`
	Title string    `json:title`
}

type Method struct {
	Id          uuid.UUID     `json:id validate:"required,uuid"`
	Title       string        `json:title`
	Description string        `json:description`
	Labels      []MethodLabel `json:labels`
}

type Content struct {
	Id    uuid.UUID `json:id validate:"required,uuid"`
	Title string    `json:title`
	Text  string    `json:text`
}
