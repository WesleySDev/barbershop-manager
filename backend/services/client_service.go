package services

import (
	"github.com/WesleySDev/barbershop-manager/database"
	"github.com/WesleySDev/barbershop-manager/models"
)

func GetClients() ([]models.Client, error) {

	var clients []models.Client

	result := database.DB.Find(&clients)

	if result.Error != nil {
		return nil, result.Error
	}
	return clients, nil

}
