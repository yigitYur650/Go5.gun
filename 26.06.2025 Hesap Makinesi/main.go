package main

import (
	"fmt"
)

func main() {
	/* var sayi1 int
	var sayi2 int
	var secim int
	var sonuc int
	fmt.Println("1-Toplama, 2-Çıkarma, 3-Çarpma, 4-Bölme")
	fmt.Println("İlk sayıyı giriniz: ")
	fmt.Scanln(&sayi1)
	fmt.Println("İkinci sayıyı giriniz: ")
	fmt.Scanln(&sayi2)
	fmt.Println("Bir secim yapınız: ")
	fmt.Scanln(&secim)
	if secim == 1 {
		sonuc = sayi1 + sayi2
		fmt.Println(sonuc)
	} else if secim == 2 {
		sonuc = sayi1 - sayi2
		fmt.Println(sonuc)
	} else if secim == 3 {
		sonuc = sayi1 * sayi2
		fmt.Println(sonuc)

	} else if secim == 4 {
		if sayi2 == 0 {
			fmt.Println("0'a bölünemez")
		} else {
			sonuc = sayi1 / sayi2
			fmt.Println(float64(sonuc))
		}

	} else {
		fmt.Println("Geçeriz işlem")
	}*/
	var sayi1 int
	var sayi2 int
	var secim int
	var sonuc float64
	fmt.Println("Önce sayıları giriniz, ardından seçim yapınız")
	fmt.Println("1-Toplama, 2-Çıkarma, 3-Çarpma, 4-Bölme")
	fmt.Print("İlk sayıyı giriniz: ")
	fmt.Scanln(&sayi1)
	fmt.Print("İkinci sayıyı giriniz: ")
	fmt.Scanln(&sayi2)
	fmt.Print("Bir seçim yapınız: ")
	fmt.Scanln(&secim)

	switch secim {
	case 1:
		sonuc = float64(sayi1 + sayi2)
		fmt.Println("Sonuç:", sonuc)
	case 2:
		sonuc = float64(sayi1 - sayi2)
		fmt.Println("Sonuç:", sonuc)
	case 3:
		sonuc = float64(sayi1 * sayi2)
		fmt.Println("Sonuç:", sonuc)
	case 4:
		if sayi2 == 0 {
			fmt.Println("Hata: 0'a bölünemez.")
		} else {
			sonuc = float64(sayi1) / float64(sayi2)
			fmt.Println("Sonuç:", sonuc)
		}
	default:
		fmt.Println("Geçersiz işlem.")
	}
} // alttaki kodun bölme kısmı daha iyi çalışıyor
