package model

import "github.com/google/uuid"

var MethodsMock = []Method{
	{
		uuid.New(),
		"Mind-Map",
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam",
		nil,
	},
	{
		uuid.New(),
		"Blitzlicht",
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam",
		nil,
	},
	{
		uuid.New(),
		"Soziometrische Abfrage",
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam",
		nil,
	},
	{
		uuid.New(),
		"Internetrecherche",
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam",
		nil,
	},
}

var BoardsMock = []Board{
	{
		uuid.New(),
		"Test Board",
	},
}
