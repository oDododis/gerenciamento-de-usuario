package service

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Estrutura do Banco de Dados

type DataBase struct {
	db *gorm.DB
}

//Cria um db para manipulação do gorm no service

func NewDB() *DataBase {
	return &DataBase{}
}

//Inicia a conecção com o Banco de Dados

func (d *DataBase) StartConnection() error {
	db, err := gorm.Open(sqlite.Open("usersDataBase.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	d.db = db
	return nil
}

// Obtém a conecção com o Banco

func (d *DataBase) GetConnection() *gorm.DB {
	return d.db
}
