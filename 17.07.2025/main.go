package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type kullanıcıYarat struct {
	Kullanıcıadı string `json:"isim"`
	Soyisim      string `json:"soyisim"`
	Email        string `json:"email"`
	Sifre        string `json:"sifre"`
	Yas          int32  `json:"age"`
}

func main() {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		fmt.Printf("Merhaba şuan middleware içindesiniz %s %s", c.BaseURL(), c.Request())
		return c.Next()

	})
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Hello world")
		return c.JSON("Hello world")
	})
	app.Get("/users/:userId", func(c *fiber.Ctx) error {
		userIDparam := c.Params("userId")
		return c.SendString(fmt.Sprintf("Senin ID bu->%s", userIDparam))
	})
	app.Post("/user", func(c *fiber.Ctx) error {
		fmt.Println("İLk endpointime hoşgeldiniz")
		var request kullanıcıYarat
		err := c.BodyParser(&request)
		if err != nil {
			return c.SendString(fmt.Sprintf("Json hatası oluştu ->%v", err.Error()))
		}
		cevapmesajı := fmt.Sprintf("Kullanıcı başarıyla oluştu %s", request.Kullanıcıadı)
		return c.Status(200).JSON(cevapmesajı)

	})
	app.Listen(":8080")
}
