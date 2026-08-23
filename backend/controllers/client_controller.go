package controllers

import (
	"strconv"

	"github.com/WesleySDev/barbershop-manager/models"
	"github.com/WesleySDev/barbershop-manager/services"
	"github.com/gofiber/fiber/v2"
)

func GetClients(c *fiber.Ctx) error {

	clients, err := services.GetClients()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(clients)
}

func CreateClient(c *fiber.Ctx) error {
	var client models.Client

	if err := c.BodyParser(&client); err != nil {
		return err
	}
	err := services.CreateClient(client)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(client)
}

func UpdateClient(c *fiber.Ctx) error {
	var client models.Client

	idString := c.Params("id")

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "ID invalido",
		})
	}
	if err = c.BodyParser(&client); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Dados invalido",
		})
	}
	err = services.UpdateClient(id, client)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(client)
}

func DeleteClient(c *fiber.Ctx) error {
	idString := c.Params("id")

	id, err := strconv.ParseInt(idString, 10, 64)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "ID invalido",
		})
	}
	err = services.DeleteClient(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})

	}
	return c.JSON(fiber.Map{
		"message": "Cliente excluido com sucesso",
	})
}
