/*
	package main

import "fmt"

	func main() {
		var isim string
		selam(isim)
	}

	func selam(isim string) {
		fmt.Println("Bir isim giriniz:  ")// burda fonksiyonu ilk kez kullandım
		fmt.Scanln(&isim)
		fmt.Println("Selam", isim)
	}
*/
/* package main

import "fmt"

func main() {
	toplam()
}
func toplam() {
	var i int
	var j int
	fmt.Println("İlk sayıyı giriniz:  ")
	fmt.Scanln(&i)
	fmt.Println("İkinci sayıyı giriniz:  ")
	fmt.Scanln(&j)
	fmt.Println("İki sayının toplamı: ", i+j)// burda fonksiyonu kullanarak kullanıcıdan aldığım değerleri toplayıp sonuç yazdırdım
} */
// package main

// import "fmt"

//	func main() {
//		sonuc := topla()
//		fmt.Println("Sonuç: ", sonuc)
//	}
//
//	func topla() int {
//		var a int
//		var b int
//		fmt.Println("Bir sayı giriniz: ")// Burda değer döndürmeyi gördüm
//		fmt.Scanln(&a)
//		fmt.Println("Bir sayı giriniz: ")
//		fmt.Scanln(&b)
//		return a + b
//	}
package main

import "fmt"

func toplaVemesajver() (int, string) {
	var a int
	var b int
	fmt.Println("Bir sayı giriniz: ")
	fmt.Scanln(&a)
	fmt.Println("Bir sayı giriniz: ")
	fmt.Scanln(&b)
	toplam := a + b
	mesaj := "Toplama işlemi başarılı"
	return toplam, mesaj
}
func main() {
	toplam, mesaj := toplaVemesajver()
	fmt.Println("Sonuç: ", toplam)
	fmt.Println("Mesaj: ", mesaj)
}
