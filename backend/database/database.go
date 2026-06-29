package database

import (
	"log"

	"github.com/WesleySDev/barbershop-manager/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=gaza dbname=barbershop port=5432 sslmode=disable "
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	} else {
		log.Println("Conectado ao PostgreSQL com sucesso!")
	}
	DB.AutoMigrate(&models.Client{})
}
