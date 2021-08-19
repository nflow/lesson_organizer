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

	if err := db.AutoMigrate(&model.Board{}, &model.Phase{}, &model.Method{}, &model.BoardGoal{}, &model.BoardPhase{}, &model.BoardMethod{}, &model.BoardContent{}); err != nil {
		return nil, err
	}

	return db, nil
}
