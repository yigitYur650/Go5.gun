package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Kitap struct {
	ID    int
	Kitap string
	Yazar string
	Oy    int
}

var kitaplar = []Kitap{
	{ID: 1, Kitap: "1984", Yazar: "George Owrell"},
	{ID: 2, Kitap: "Hayvan Çiftliği", Yazar: "George Owrell"},
}

func TokenKontrol(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token != "Gizli token" {
		return c.Status(401).JSON(fiber.Map{
			"error": "Token hatalı",
		})
	}
	return c.Next()
}

func main() {
	app := fiber.New()
	TokenKontrol(c * fiber.Ctx)
	app.Post("/kitap/:id/oy", func(c *fiber.Ctx) error {
		id := c.Params("id")
		num, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(400).SendString("Geçersiz ID")
		}
		for i, k := range kitaplar {
			if k.ID == num {
				kitaplar[i].Oy++
				return c.JSON(fiber.Map{
					"message": "Oy verildi",
				})
			}
		}
		return c.Status(400).SendString("Kitap bulunamadı")
	})
	app.Get("/kitap/:id/oy", func(c *fiber.Ctx) error {
		id := c.Params("id")
		num, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(400).SendString("Geçersiz ID")
		}
		for _, k := range kitaplar {
			if k.ID == num {
				return c.JSON(fiber.Map{
					"kitap": k.Kitap,
					"oy":    k.Oy,
				})
			}
		}
		return c.Status(400).SendString("Kitap bulunamadı")
	})
	app.Listen(":8080")
}
