package main

import (
	"JWT_auth/go-auth/database"
	"JWT_auth/go-auth/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("Hello, World!")
	})

	app.Get("/view", func(c *fiber.Ctx) error {
		return c.SendString("view Hello, World!")
	})

	app.Post("/register", service.Register)

	app.Post("/login", service.Login)

	app.Post("/logout", service.Logout)

	app.Get("/user/:id", service.User)

	app.Listen(":8000")
}
