package main

import (
	"github.com/gofiber/fiber/v2"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Welcome to Jason's Coding Dojo for Ninjas")
}

func Routers(app *fiber.App) {
	app.Get("/players", player.GetPlayers)
	app.Get("/player/:id", player.GetPlayer)
	app.Post("/player", player.SavePlayer)
	app.Delete("/player/:id", player.DeletePlayer)
	app.Put("/player/:id", player.UpdatePlayer)
}

func main() {
	player.InitialMigration()
	app := fiber.New()
	app.Get("/", hello)
	Routers(app)
	app.Listen(":3000")
}
