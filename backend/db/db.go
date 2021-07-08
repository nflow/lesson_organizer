package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.SetupJoinTable(&Board{}, "Contents", &BoardContents{}); err != nil {
		return nil, err
	}

	if err := db.SetupJoinTable(&Board{}, "Phases", &BoardPhases{}); err != nil {
		return nil, err
	}

	if err := db.SetupJoinTable(&Phase{}, "Methods", &PhaseMethods{}); err != nil {
		return nil, err
	}

	if err := db.SetupJoinTable(&PhaseMethods{}, "Contents", &MethodContents{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&Board{}, &Content{}, &Phase{}, &Method{}); err != nil {
		return nil, err
	}

	return db, nil
}
