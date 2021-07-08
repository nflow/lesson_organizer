package main

import (
	"github.com/nflow/lesson_organizer/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.SetupJoinTable(&model.Board{}, "Phases", &model.BoardPhases{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&model.Board{}, &model.Phase{}, &model.Method{}, &model.Content{}); err != nil {
		return err
	}

	return nil
}
