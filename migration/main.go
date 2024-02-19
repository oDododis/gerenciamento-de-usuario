package main

import (
	"Teste/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

// Crias as tabelas do Banco vazias

func main() {
	db, err := gorm.Open(sqlite.Open("usersDataBase.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err = db.AutoMigrate(&model.User{}, &model.Token{}); err != nil {
		log.Fatal(err)
	}
}
