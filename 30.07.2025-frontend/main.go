package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"golang.org/x/crypto/bcrypt"
)

// Kullanıcı verileri
type kullanıcıYarat struct {
	Kullanıcıadı string `json:"isim" validate:"required,min=2"`
	Soyisim      string `json:"soyisim" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Sifre        string `json:"sifre" validate:"required,min=2,max=16"`
	Yas          int32  `json:"age" validate:"required,gte=0,lte=150"`
}

// Validasyon hata yapısı
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

// Login için struct
type GirisCevabı struct {
	Email string `json:"email"`
	Sifre string `json:"sifre"`
}

// Kayıtlı kullanıcılar (memory)
var kullanıcılarMap = make(map[string]kullanıcıYarat)

// Validasyon fonksiyonu
func validate(data interface{}) []customErrorValidation {
	var errorsList []customErrorValidation
	validate := validator.New()

	if err := validate.Struct(data); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			errorsList = append(errorsList, customErrorValidation{
				Hata:   true,
				Alan:   fieldErr.StructField(),
				Etiket: fieldErr.Tag(),
				Param:  fieldErr.Param(),
				Value:  fieldErr.Value(),
			})
		}
	}
	return errorsList
}

func main() {
	app := fiber.New()

	// Panic önleyici middleware
	app.Use(recover.New())

	// Statik dosyaları ./public klasöründen sun
	app.Static("/", "./public")

	// Kullanıcı kaydı (form urlencoded olarak gönderilecek)
	app.Post("/user", func(c *fiber.Ctx) error {
		// Form verisi al
		isim := c.FormValue("isim")
		soyisim := c.FormValue("soyisim")
		email := c.FormValue("email")
		sifre := c.FormValue("sifre")
		ageStr := c.FormValue("age")

		age, err := strconv.Atoi(ageStr)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Yaş geçerli sayı değil"})
		}

		kullanici := kullanıcıYarat{
			Kullanıcıadı: isim,
			Soyisim:      soyisim,
			Email:        email,
			Sifre:        sifre,
			Yas:          int32(age),
		}

		// Validasyon
		if errors := validate(kullanici); errors != nil && errors[0].Hata {
			var hataYanıt HataYanıtı
			hataYanıt.Status = http.StatusBadRequest
			for _, err := range errors {
				hataYanıt.HataDetayları = append(hataYanıt.HataDetayları, HataYanıtıDetayları{
					Description: fmt.Sprintf("Alan %s için '%s' kuralı ihlal edildi", err.Alan, err.Etiket),
					AlanHatası:  err.Alan,
				})
			}
			return c.Status(http.StatusBadRequest).JSON(hataYanıt)
		}

		// Şifreyi hashle
		hash, err := bcrypt.GenerateFromPassword([]byte(kullanici.Sifre), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Şifre hashlenirken hata oluştu"})
		}
		kullanici.Sifre = string(hash)

		// Kaydet (email bazlı)
		kullanıcılarMap[kullanici.Email] = kullanici

		return c.SendString(fmt.Sprintf("Kullanıcı başarıyla oluştu: %s", kullanici.Kullanıcıadı))
	})

	// Login endpointi (JSON bekliyor)
	app.Post("/login", func(c *fiber.Ctx) error {
		var req GirisCevabı
		if err := c.BodyParser(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz JSON formatı"})
		}

		kullanici, exists := kullanıcılarMap[req.Email]
		if !exists {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Kullanıcı bulunamadı"})
		}

		err := bcrypt.CompareHashAndPassword([]byte(kullanici.Sifre), []byte(req.Sifre))
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Şifre yanlış"})
		}

		return c.JSON(fiber.Map{"message": "Giriş başarılı"})
	})

	fmt.Println("Sunucu 8080 portunda başlatıldı")
	app.Listen(":8080")
}
