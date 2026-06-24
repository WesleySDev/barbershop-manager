package controllers

import (
	"github.com/WesleySDev/barbershop-manager/services"
	"github.com/gofiber/fiber/v2"
)

func GetClients(c *fiber.Ctx) error {

	clients := services.GetClients()

	return c.JSON(clients)
}
