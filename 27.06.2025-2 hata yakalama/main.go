package main

import "fmt"

func main() {
	bolme()
}
func bolme() {
	defer hatayakala() // defer fonksiyonu burada olmalı
	var a float64
	var b float64
	fmt.Println("Bölmek istediğiniz sayıyı giriniz: ")
	fmt.Scanln(&a)
	fmt.Println("Kaça bölecekseniz o sayıyı giriniz: ")
	fmt.Scanln(&b)
	if b == 0 {
		panic("Bir sayı 0'a bölünemez") // bir hata mesajı fırlatıyor
	}
	sonuc := a / b
	fmt.Println("Sonuç", sonuc)
}
func hatayakala() {
	if r := recover(); r != nil { // nil boş değer dönüyor // recoverde hata mesajını alıp gerçekten hata mı diye nil fonksiyonunda kontrol ediyor eğer hata varsa if çalışıyor
		fmt.Println("Hata yakalandı", r) // buraya r kotmazsak bolme fonksiyonundaki if b==0 daki mesajı göremeyiz
	}
}
