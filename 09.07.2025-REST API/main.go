package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Kitap struct {
	ID    int    `json:"id"`
	Kitap string `json:"kitap"`
	Yazar string `json:"yazar"`
}

var kitaplar = []Kitap{
	{ID: 1, Kitap: "1984", Yazar: "George Orwell"},
	{ID: 2, Kitap: "Hayvan Çiftliği", Yazar: "George Orwell"},
}

func main() {
	app := fiber.New()

	// POST /kitaplar
	app.Post("/kitaplar", func(c *fiber.Ctx) error { //Kitap ekleniyor
		t := new(Kitap)
		if err := c.BodyParser(t); err != nil {
			return c.Status(400).SendString("Hatalı veri")
		}
		t.ID = len(kitaplar) + 1
		kitaplar = append(kitaplar, *t)
		return c.JSON(fiber.Map{
			"message": "Kitap başarıyla eklendi",
			"kitap":   t,
		})
	})

	app.Get("/kitaplar", func(c *fiber.Ctx) error { //Kitapları gösterir
		return c.JSON(kitaplar)
	})
	app.Get("/kitaplar/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		num, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(400).SendString("ID'yi bulamadık")
		}
		for _, k := range kitaplar {
			if k.ID == num {
				return c.JSON(k)
			}
		}
		return c.Status(400).SendString("Kitabı bulamadık")
	})
	app.Delete("/kitaplar/:id", func(c *fiber.Ctx) error { //Kitapları siler
		id := c.Params("id")
		num, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(400).SendString("ID'yi bulamadık")
		}
		for i, k := range kitaplar {
			if k.ID == num {
				kitaplar = append(kitaplar[:i], kitaplar[i+1:]...)
				return c.SendStatus(204)
			}
		}
		return c.Status(400).SendString("Kitabı silemedik")
	})
	app.Put("/kitaplar/:id", func(c *fiber.Ctx) error { // komple kitapları günceller
		id := c.Params("id")
		num, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(400).SendString("ID bulunamadı")
		}
		t := new(Kitap)
		if err := c.BodyParser(t); err != nil {
			return c.Status(400).SendString("Geçersiz veri")
		}
		for i, k := range kitaplar {
			if k.ID == num {
				t.ID = num
				kitaplar[i] = *t
				return c.JSON(t)
			}
		}
		return c.Status(400).SendString("Kitap bulunamadı")
	})
	app.Patch("/kitaplar/:id", func(c *fiber.Ctx) error { // kitabın belli bir kısmını günceller
		id := c.Params("id")
		num, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(400).SendString("Geçersiz ID")
		}
		t := new(Kitap)
		if err := c.BodyParser(t); err != nil {
			return c.Status(400).SendString("Geçersiz veri")
		}
		for i, k := range kitaplar {
			if k.ID == num {
				if t.Kitap != "" {
					kitaplar[i].Kitap = t.Kitap
				}
				if t.Yazar != "" {
					kitaplar[i].Yazar = t.Yazar
				}
				return c.JSON(kitaplar[i])
			}
		}
		return c.Status(404).SendString("Kitap bulunamadı")
	})

	app.Listen(":8080")
}
