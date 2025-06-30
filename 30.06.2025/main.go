package main

import "fmt"

func main() {
	var sayı []int
	var girileceksayi int
	fmt.Println("Kaç sayı girileceğini yazınız: ")
	fmt.Scanln(&girileceksayi)
	fmt.Printf("%d adet sayı giriniz\n", girileceksayi)
	for i := 0; i < girileceksayi; i++ {
		fmt.Scanln(&sayı)
		var sayıekle int // Burada slince e direkt değer atayamadığımızdan dolayı bir değer oluşturuyoruz alacağımız sayıları buna atıp burdanda append fonksiyonuyla slince e aktarıyoruz
		fmt.Scanln(&sayıekle)
		sayı = append(sayı, sayıekle)
	}
	fmt.Println("Dizi içerisindeki sayılar:\n", sayı)
}
