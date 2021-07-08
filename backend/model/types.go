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
	Id       uuid.UUID `json:id validate:"required,uuid" grom:"primaryKey"`
	Name     string    `json:name`
	Goals    []Goal    `json:goals`
	Contents []Content `json:contents`
}

type Goal struct {
	Id   uuid.UUID `json:id validate:"required,uuid" grom:"primaryKey"`
	Text string    `json:text`
}

type Phase struct {
	Id      uuid.UUID `json:id validate:"required,uuid" grom:"primaryKey"`
	Title   string    `json:title`
	Methods []Method  `json:methods`
}

type Method struct {
	Id          uuid.UUID     `json:id validate:"required,uuid" grom:"primaryKey"`
	Title       string        `json:title`
	Description string        `json:description`
	Labels      []MethodLabel `json:labels`
	Contents    []Content     `json:contents`
}

type Content struct {
	Id    uuid.UUID `json:id validate:"required,uuid" grom:"primaryKey"`
	Title string    `json:title`
	Text  string    `json:text`
}
