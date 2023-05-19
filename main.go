package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mergen1212/fiber-api-go/datebase"
	"github.com/mergen1212/fiber-api-go/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("welcome to my API")
}
func setupRouter(app *fiber.App) {
	app.Get("/api", welcome)
	app.Post("api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
}
func main() {
	app := fiber.New()
	datebase.ConnectDB()
	setupRouter(app)

	log.Fatal(app.Listen(":3000"))
}
