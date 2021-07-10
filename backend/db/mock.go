package db

import (
	"github.com/google/uuid"
	"github.com/nflow/lesson_organizer/model"
)

var MethodsMock = []model.Method{
	{
		uuid.New(),
		"Mind-Map",
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam",
		[]model.MethodLabel{},
	},
	{
		uuid.New(),
		"Blitzlicht",
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam",
		[]model.MethodLabel{},
	},
	{
		uuid.New(),
		"Soziometrische Abfrage",
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam",
		[]model.MethodLabel{},
	},
	{
		uuid.New(),
		"Internetrecherche",
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam",
		[]model.MethodLabel{},
	},
}

var BoardsMock = []model.Board{
	{
		uuid.New(),
		"Test Board",
		nil,
		nil,
		nil,
	},
}
