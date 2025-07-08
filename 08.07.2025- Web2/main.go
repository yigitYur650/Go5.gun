package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OrkunYuvaya") // ana menüde orkun yuvaya yazdırır
	})
	type Todo struct {
		Name string `json:"name"`
	}
	app.Get("/user/:id", func(c *fiber.Ctx) error { // URL'den gelen kullanıcı ID'sini yakalar ve ekrana yazdırır

		id := c.Params("id")
		return c.SendString("Kullanıcı ID:" + id)
	})
	app.Post("/new", func(c *fiber.Ctx) error { //Buradada postman üzerinden string atayabiliyoruz
		t := new(Todo)
		if err := c.BodyParser(t); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Kötü istek", // Çalışmazsa burası çalışır
			})
		}
		return c.JSON(fiber.Map{
			"message": "Todo başarıyla oluşturuldu", // Postman başarıyla atarsa burası çalışır
			"name":    t.Name,
		})
	})
	app.Listen(":8080")
}
