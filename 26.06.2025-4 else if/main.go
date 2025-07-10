package main

import "fmt"

func main() {
	/* var sayi int
	fmt.Println("Bir sayı giriniz: ")
	fmt.Scanln(&sayi)
	if sayi < 0 {
		fmt.Println("Sayı negatif")
	} else if sayi == 0 {
		fmt.Println("Sayı 0' a eşit")
	} else {
		fmt.Println("Sayı pozitif")
	} */
	var not int
	fmt.Println("Notunuzu giriniz")
	fmt.Scanln(&not)
	if not > 85 { // aralıkları bilerek böyle yazdım 85'e tebrikler çok başarılısın demek gelmedi içimden
		fmt.Println("Tebrikler çok başarılısın")
	} else if not > 70 {
		fmt.Println("Tebrik ederim")
	} else if not > 50 {
		fmt.Println("Dersi geçtin")
	} else {
		fmt.Println("Kaldın")
	}
}
