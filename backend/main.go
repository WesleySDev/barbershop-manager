package main

import (
	"github.com/WesleySDev/barbershop-manager/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New() // Criando um sevidor web novo e armazenando na variavel "app"

	routes.SetupRoutes(app)

	app.Listen(":3000") // ligação do sevidor (Comece a ouvir requisições na porta 3000.)
}
