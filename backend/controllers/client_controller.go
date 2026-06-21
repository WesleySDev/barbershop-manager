package controllers

import (
	"github.com/WesleySDev/barbershop-manager/models"
	"github.com/gofiber/fiber/v2"
)

func GetClients(c *fiber.Ctx) error {

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

	return c.JSON(clients)
}
