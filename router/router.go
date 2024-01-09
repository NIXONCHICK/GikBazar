package router

import (
	"GikBazar/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/user", controller.SetUser)
	app.Get("/user", controller.GetUsers)
	app.Get("/user:id", controller.GetUserById)

	app.Post("/shop", controller.PostItem)
	app.Get("/shop", controller.GetItems)
	app.Get("/shop:id", controller.GetItemById)
	app.Delete("/shop:id", controller.DeleteItem)
	app.Post("shop:id", controller.UpdateItem)
	
	app.Post("/auth", controller.Authorisation)
	app.Get("/cart:id", controller.ShoperById)
	app.Post("/cart:id", controller.UpdateShoper)
	app.Post("/cart", controller.SetItemToShoper)
}