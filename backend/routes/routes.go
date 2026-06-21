package routes

import (
	"github.com/WesleySDev/barbershop-manager/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controllers.GetHome) // Rota para a página inicial, quando o usuário acessar a raiz do site, ele vai chamar a função GetHome do controllers.
	app.Get("/clientes", controllers.GetClients)
}
