package main

import (
	"fiber_test/dal"
	"fiber_test/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect() // burdada bunu çalıştırmamızın sebebi defalarca db çağırmadan halledelim diye
	database.DB.AutoMigrate(&dal.Todo{})
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OrkunYuvaya")
	})
	app.Listen(":8080")
}
