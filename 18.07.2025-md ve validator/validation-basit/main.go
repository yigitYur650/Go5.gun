package main

import (
	"fmt"
	"strings"
)

func main() {
	isim := "Yiğit"
	email := "yigit@gmail.com"
	yas := 17

	if isim == "" {
		fmt.Println("İsim boş olamaz.")
	}

	if !strings.Contains(email, "@") {
		fmt.Println("Geçersiz email adresi.")
	}

	if yas < 18 {
		fmt.Println("Yaş 18 veya daha büyük olmalı.")
	}

	fmt.Println("Doğrulama tamamlandı.")
}
