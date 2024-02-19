package service

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type dataBase struct {
	db *gorm.DB
}

func (d *dataBase) StartConnection() error {
	db, err := gorm.Open(sqlite.Open("usersDataBase.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	d.db = db
	return nil
}

func (d *dataBase) GetConnection() *gorm.DB {
	return d.db
}
