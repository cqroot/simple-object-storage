package models

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Object struct {
	Name      string `gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ListObjects(dsn string) ([]Object, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	var objects []Object
	err = db.Select("name").Find(&objects).Error
	if err != nil {
		return nil, err
	}

	return objects, err
}

func PutObject(dsn string, name string) error {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&Object{})

	err = db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&Object{
		Name: name,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func DeleteObject(dsn string, name string) error {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.Where("name = ?", name).Delete(&Object{}).Error
	if err != nil {
		return err
	}

	return nil
}
