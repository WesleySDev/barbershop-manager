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

func CreateClient(client models.Client) error {
	result := database.DB.Create(&client) // insert no postgres
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateClient(id int64, client models.Client) error {
	var existingClient models.Client
	result := database.DB.First(&existingClient, id)\

	if result.Error != nil {
		return result.Error
	}
	existingClient.Name = client.Name
	existingClient.Telephone = client.Telephone
	existingClient.Email = client.Email

	if result := database.DB.Save(&existingClient);
	result.Error != nil {
		return result.Error
	}
	return nil
}
