package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/uuid"
)

type kullanıcıYarat struct {
	Kullanıcıadı string `json:"isim" validate:"required,min=2"`
	Soyisim      string `json:"soyisim" validate:"required"`
	Email        string `json:"email" validate:"required"`
	Sifre        string `json:"sifre" validate:"required,min=2,max=16"`
	Yas          int32  `json:"age" validate:"required"`
}

type HataYanıtı struct {
	Status        int32                 `json:"status"`
	HataDetayları []HataYanıtıDetayları `json:"HataDetayları"`
}

type HataYanıtıDetayları struct {
	Description string `json:"description"`
	AlanHatası  string `json:"AlanHatası"`
}

type customErrorValidation struct {
	Hata   bool
	Alan   string
	Etiket string
	Param  string
	Value  interface{}
}

func validate(data interface{}) []customErrorValidation {
	var errorsList []customErrorValidation
	var validate = validator.New()
	if err := validate.Struct(data); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			var cve customErrorValidation
			cve.Hata = true
			cve.Alan = fieldErr.StructField()
			cve.Etiket = fieldErr.Tag()
			cve.Param = fieldErr.Param()
			cve.Value = fieldErr.Value()
			errorsList = append(errorsList, cve)
		}
	}
	return errorsList
}

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		fmt.Printf("Merhaba şuan middleware içindesiniz %s %s", c.BaseURL(), c.Request())
		return c.Next()
	})

	app.Use(recover.New())

	app.Use("/user", func(c *fiber.Ctx) error {
		collerationID := c.Get("x-collerationID")
		if collerationID == "" {
			return c.Status(fiber.ErrBadRequest.Code).JSON("Boş bir bağlantı kodu girdiniz.")
		}
		_, err := uuid.Parse(collerationID)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON("ID GUID formatında değil")
		}
		c.Locals("collerationID", collerationID)
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
		fmt.Println("İlk endpointime hoşgeldiniz")

		var request kullanıcıYarat
		err := c.BodyParser(&request)
		if err != nil {
			return c.SendString(fmt.Sprintf("Json hatası oluştu ->%v", err.Error()))
		}

		if error := validate(request); error != nil && error[0].Hata {
			var AlanHatası HataYanıtı
			var HataDetayları HataYanıtıDetayları
		}

		cevapmesajı := fmt.Sprintf("Kullanıcı başarıyla oluştu %s", request.Kullanıcıadı)
		return c.Status(200).JSON(cevapmesajı)
	})

	app.Listen(":8080")
}
