package main

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

type Kitap struct {
	ID    int    `json:"id"`
	Kitap string `json:"kitap"`
	Yazar string `json:"yazar"`
}

func main() {
	app := fiber.New()
	db, err := sql.Open("sqlite3", "kitaplar.db")
	if err != nil {
		log.Fatal("Veritabanına bağlanılamadı. Neden olan hata:", err)
	}
	defer db.Close()
	log.Println("Veritabaına başarıyla bağlanıldı")

	// Tablo oluştur (varsa)
	createTableSQL := `CREATE TABLE IF NOT EXISTS kitaplar (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		kitap TEXT,
		yazar TEXT
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Tablo oluşturma hatası:", err)
	}

	// POST /kitaplar
	app.Post("/kitaplar", func(c *fiber.Ctx) error {
		t := new(Kitap)
		if err := c.BodyParser(t); err != nil {
			return c.Status(400).SendString("Hatalı veri")
		}
		stmt, err := db.Prepare("INSERT INTO kitaplar (kitap,yazar) VALUES(?,?)")
		if err != nil {
			return c.Status(500).SendString("Hazırlanamadı")
		}
		defer stmt.Close()

		res, err := stmt.Exec(t.Kitap, t.Yazar)
		if err != nil {
			return c.Status(500).SendString("Veritabanına yazılamadı")
		}

		id, err := res.LastInsertId()
		if err == nil {
			t.ID = int(id)
		}

		return c.JSON(fiber.Map{
			"message": "Kitap başarıyla eklendi",
			"kitap":   t,
		})
	})

	// GET /kitaplar
	app.Get("/kitaplar", func(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT id,kitap,yazar FROM kitaplar")
		if err != nil {
			return c.Status(500).SendString("Veri alınamadı")
		}
		defer rows.Close()

		var gelenkitaplar []Kitap
		for rows.Next() {
			var k Kitap
			if err := rows.Scan(&k.ID, &k.Kitap, &k.Yazar); err != nil {
				continue
			}
			gelenkitaplar = append(gelenkitaplar, k)
		}
		return c.JSON(gelenkitaplar)
	})

	// GET /kitaplar/:id
	app.Get("/kitaplar/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		num, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(400).SendString("ID'yi bulamadık")
		}

		row := db.QueryRow("SELECT id,kitap,yazar FROM kitaplar WHERE id = ?", num)
		var k Kitap
		err = row.Scan(&k.ID, &k.Kitap, &k.Yazar)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(404).SendString("Kitap bulunamadı")
			}
			return c.Status(500).SendString("Veritabanı hatası")
		}
		return c.JSON(k)
	})

	// PUT /kitaplar/:id
	app.Put("/kitaplar/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		num, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(400).SendString("ID bulunamadı")
		}
		t := new(Kitap)
		if err := c.BodyParser(t); err != nil {
			return c.Status(400).SendString("Geçersiz veri")
		}

		res, err := db.Exec("UPDATE kitaplar SET kitap = ?, yazar = ? WHERE id = ?", t.Kitap, t.Yazar, num)
		if err != nil {
			return c.Status(500).SendString("Güncelleme hatası")
		}
		affected, err := res.RowsAffected()
		if err != nil || affected == 0 {
			return c.Status(404).SendString("Kitap bulunamadı")
		}

		t.ID = num
		return c.JSON(t)
	})

	// PATCH /kitaplar/:id
	app.Patch("/kitaplar/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		num, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(400).SendString("Geçersiz ID")
		}
		t := new(Kitap)
		if err := c.BodyParser(t); err != nil {
			return c.Status(400).SendString("Geçersiz veri")
		}

		row := db.QueryRow("SELECT kitap,yazar FROM kitaplar WHERE id = ?", num)
		var mevcut Kitap
		err = row.Scan(&mevcut.Kitap, &mevcut.Yazar)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(404).SendString("Kitap bulunamadı")
			}
			return c.Status(500).SendString("Veritabanı hatası")
		}

		if t.Kitap == "" {
			t.Kitap = mevcut.Kitap
		}
		if t.Yazar == "" {
			t.Yazar = mevcut.Yazar
		}

		_, err = db.Exec("UPDATE kitaplar SET kitap = ?, yazar = ? WHERE id = ?", t.Kitap, t.Yazar, num)
		if err != nil {
			return c.Status(500).SendString("Güncelleme hatası")
		}

		t.ID = num
		return c.JSON(t)
	})

	// DELETE /kitaplar/:id
	app.Delete("/kitaplar/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		num, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(400).SendString("Geçersiz ID")
		}

		res, err := db.Exec("DELETE FROM kitaplar WHERE id = ?", num)
		if err != nil {
			return c.Status(500).SendString("Silme hatası")
		}
		affected, err := res.RowsAffected()
		if err != nil || affected == 0 {
			return c.Status(404).SendString("Kitap bulunamadı")
		}
		return c.SendStatus(204)
	})

	log.Fatal(app.Listen(":8080"))
}
