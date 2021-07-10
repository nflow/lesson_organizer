package db

import (
	"github.com/nflow/lesson_organizer/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("local.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err := db.SetupJoinTable(&model.Board{}, "Contents", &BoardContents{}); err != nil {
		return nil, err
	}

	if err := db.SetupJoinTable(&model.Board{}, "Phases", &BoardPhases{}); err != nil {
		return nil, err
	}

	if err := db.SetupJoinTable(&model.Phase{}, "Methods", &PhaseMethods{}); err != nil {
		return nil, err
	}

	if err := db.SetupJoinTable(&PhaseMethods{}, "Contents", &MethodContents{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&model.Board{}, &model.Content{}, &model.Phase{}, &model.Method{}); err != nil {
		return nil, err
	}

	return db, nil
}
