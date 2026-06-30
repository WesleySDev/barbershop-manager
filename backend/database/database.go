package database

import (
	"log"

	"github.com/WesleySDev/barbershop-manager/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=gaza dbname=barbershop port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}
	log.Println("Conectado ao PostgreSQL com sucesso!")

	if err := DB.AutoMigrate(&models.Client{}); err != nil {
		log.Fatal("Erro ao criar tabela:", err)
	}
	log.Println("Tabela clients sincronizada com sucesso!")

}
