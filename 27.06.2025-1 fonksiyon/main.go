/*
	package main

import "fmt"

	func main() {
		var fakt int
		var sonuc int = 1
		fmt.Println("Faktöriyelini hesaplamak istediğiniz sayıyı giriniz: ")
		fmt.Scanln(&fakt)
		if fakt < 0 {
			fmt.Println("Negatif sayı girdiniz. Lütfen pozitif bir sayı giriniz") // Burada fonksiyon kullanmadan Faktöriyel hesabı yat-ptım
			return
		}
		for i := 1; i <= fakt; i++ {
			sonuc *= i
		}
		fmt.Println("Sonuç: ", sonuc)
	}
*/
package main

import "fmt"

func main() {
	fakt()
}
func fakt() int {
	var a int
	var sonuc int = 1
	fmt.Println("Faktöriyelini hesaplamak istediğiniz sayıyı giriniz: ")
	fmt.Scanln(&a)
	if a < 0 {
		fmt.Println("Negatif bir sayı girdiniz. Lütfen pozitif bir değer giriniz.") // Burada ise fonksiyon kullanarak yaptım
		return 0
	}
	for i := 1; i <= a; i++ {
		sonuc *= i
	}
	fmt.Println("Sonuç: ", sonuc)
	return sonuc
}
/* package main */
/*  */
/* import "fmt" */
/*  */
/* func power(x, n int) int { */
/* 	if n == 0 { */
/* 		return 1 */
/* 	} */
/* 	return x * power(x, n-1) */
/* } */
/*  */
/* func main() { */
/* 	var taban, us int */
/* 	fmt.Println("Taban sayıyı giriniz:") */
/* 	fmt.Scanln(&taban) */
/* 	fmt.Println("Üs sayıyı giriniz:") */
/* 	fmt.Scanln(&us) */
/* 	fmt.Println("Sonuç:", power(taban, us)) */
/* } */
/*  */
