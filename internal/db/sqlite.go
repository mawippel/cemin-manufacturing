package db

import (
	"controleConfeccao/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("./controleConfeccao.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto migrate the schema
	err = DB.AutoMigrate(&models.Colaborador{}, &models.Modelo{}, &models.Operacao{}, &models.Execucao{})
	if err != nil {
		log.Fatal(err)
	}
}
