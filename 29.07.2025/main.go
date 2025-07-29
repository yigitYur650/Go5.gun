package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"         // Girdi validasyonu için kullanılır
	"github.com/gofiber/fiber/v2"                    // Fiber web framework'ü
	"github.com/gofiber/fiber/v2/middleware/recover" // Panic yakalama için middleware
	"github.com/google/uuid"                         // UUID doğrulama için paket
	"golang.org/x/crypto/bcrypt"                     // Şifre hashleme için bcrypt
)

// Kullanıcı verilerini tutan struct
type kullanıcıYarat struct {
	Kullanıcıadı string `json:"isim" validate:"required,min=2"`         // En az 2 karakter zorunlu
	Soyisim      string `json:"soyisim" validate:"required"`            // Zorunlu alan
	Email        string `json:"email" validate:"required,email"`        // Geçerli email zorunlu
	Sifre        string `json:"sifre" validate:"required,min=2,max=16"` // 2-16 karakter arası zorunlu
	Yas          int32  `json:"age" validate:"required,gte=0,lte=150"`  // 0-150 arası yaş
}

// Validasyon hatalarını kullanıcılara göstermek için struct
type HataYanıtı struct {
	Status        int32                 `json:"status"`        // HTTP durum kodu
	HataDetayları []HataYanıtıDetayları `json:"HataDetayları"` // Her bir alan hatasının detayları
}

// Tek bir alan hatası için detay struct'ı
type HataYanıtıDetayları struct {
	Description string `json:"description"` // Açıklama
	AlanHatası  string `json:"AlanHatası"`  // Hatalı alan adı
}

// Validasyon sırasında dönecek iç yapı (dahili)
type customErrorValidation struct {
	Hata   bool
	Alan   string
	Etiket string
	Param  string
	Value  interface{}
}

// Login isteği için struct
type GirisCevabı struct {
	Email string `json:"email"` // Kullanıcının emaili
	Sifre string `json:"sifre"` // Kullanıcının düz metin şifresi
}

// Globalde kullanıcıları saklamak için map (email bazlı)
var kullanıcılarMap = make(map[string]kullanıcıYarat)

// Validasyon fonksiyonu: struct'ı kontrol eder ve hata döner
func validate(data interface{}) []customErrorValidation {
	var errorsList []customErrorValidation
	var validate = validator.New()

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

	// Middleware: Her gelen istekte çalışır, log basar
	app.Use(func(c *fiber.Ctx) error {
		fmt.Printf("Merhaba şuan middleware içindesiniz %s %s\n", c.BaseURL(), c.Request())
		return c.Next()
	})

	// Panic durumunda uygulamanın çökmesini engeller
	app.Use(recover.New())

	// "/user" altındaki endpoint'ler için collerationID kontrolü
	app.Use("/user", func(c *fiber.Ctx) error {
		collerationID := c.Get("x-collerationID")
		if collerationID == "" {
			return c.Status(fiber.ErrBadRequest.Code).JSON("Boş bir bağlantı kodu girdiniz.")
		}
		_, err := uuid.Parse(collerationID)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON("ID GUID formatında değil")
		}
		c.Locals("collerationID", collerationID) // Sonraki handler'lara aktarım
		return c.Next()
	})

	// Basit bir GET endpointi
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Hello world")
		return c.JSON("Hello world")
	})

	// Parametreli endpoint örneği
	app.Get("/users/:userId", func(c *fiber.Ctx) error {
		userIDparam := c.Params("userId")
		return c.SendString(fmt.Sprintf("Senin ID bu->%s", userIDparam))
	})

	// Login endpointi
	app.Post("/login", func(c *fiber.Ctx) error {
		var req GirisCevabı
		if err := c.BodyParser(&req); err != nil {
			//Bodyparse edilemezse 400 dönecek
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Geçersiz JSON formatı",
			})
		}
		// Kullanıcının kayıt olup olmadığını kontrol et
		kullanici, exists := kullanıcılarMap[req.Email]
		if !exists {
			// Eğer kullanıcı yoksa 404 döner
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Kullanıcı bulunamadı",
			})
		}
		// Hashlenmiş şifre ile gelen şifreyi karşılaştırır
		err := bcrypt.CompareHashAndPassword([]byte(kullanici.Sifre), []byte(req.Sifre))
		if err != nil {
			// Şifre uyuşmazsa hata döner
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Şifre yanlış",
			})
		}
		// Her şey doğruysa başarılı mesajı döner
		return c.JSON(fiber.Map{
			"message": "Giriş başarılı",
		})
	})

	// Kullanıcı oluşturma endpointi
	app.Post("/user", func(c *fiber.Ctx) error {
		fmt.Println("İlk endpointime hoşgeldiniz")

		var request kullanıcıYarat

		// Gelen JSON'u struct'a parse et
		err := c.BodyParser(&request)
		if err != nil {
			return c.SendString(fmt.Sprintf("Json hatası oluştu ->%v", err.Error()))
		}

		// Validasyon kontrolü
		if errors := validate(request); errors != nil && errors[0].Hata {
			var hataYanıt HataYanıtı
			hataYanıt.Status = http.StatusBadRequest

			// Tüm validasyon hatalarını kullanıcıya döndür
			for _, err := range errors {
				hataYanıt.HataDetayları = append(hataYanıt.HataDetayları, HataYanıtıDetayları{
					Description: fmt.Sprintf("Alan %s için '%s' kuralı ihlal edildi", err.Alan, err.Etiket),
					AlanHatası:  err.Alan,
				})
			}
			return c.Status(http.StatusBadRequest).JSON(hataYanıt)
		}

		// Şifreyi hashle
		hash, err := bcrypt.GenerateFromPassword([]byte(request.Sifre), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Şifre hashlenirken hata oluştu"})
		}
		request.Sifre = string(hash) // Hashlenmiş şifreyi kaydet

		// Kullanıcıyı map'e kaydet (email bazlı)
		kullanıcılarMap[request.Email] = request

		cevapmesajı := fmt.Sprintf("Kullanıcı başarıyla oluştu %s", request.Kullanıcıadı)
		return c.Status(200).JSON(cevapmesajı)
	})

	// Sunucuyu başlat
	app.Listen(":8080")
}
