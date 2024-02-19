package service

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DataBase struct {
	db *gorm.DB
}

func NewDB() *DataBase {
	return &DataBase{}
}

func (d *DataBase) StartConnection() error {
	db, err := gorm.Open(sqlite.Open("usersDataBase.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	d.db = db
	return nil
}

func (d *DataBase) GetConnection() *gorm.DB {
	return d.db
}
