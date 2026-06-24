package services

import (
	"github.com/WesleySDev/barbershop-manager/models"
)

func GetClients() []models.Client {

	clients := []models.Client{
		{
			ID:   1,
			Name: "Wesley",
		},
		{
			ID:   2,
			Name: "Pedro",
		},
	}
	return clients
}
