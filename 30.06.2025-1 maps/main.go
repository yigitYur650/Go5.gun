package main

import (
	"fmt"
)

func main() {
	futbolcu := make(map[string]int)
	var ogr int
	fmt.Println("Kaç Futbolcunun lisansını çıkartacaksınız :  ")
	fmt.Scanln(&ogr)
	for i := 0; i < ogr; i++ {
		var isim string
		var numara int
		fmt.Println("******Önemli uyarı******\nFutbolcu ismi ve soyismini birleşik giriniz") //burda bufio yu bilmediğim için basite kaçtım ama onuda öğreneceğim
		fmt.Println("Futbolcu ismini giriniz: ")
		fmt.Scanln(&isim)
		fmt.Println("Forma numarasını giriniz: ")
		fmt.Scanln(&numara)

		futbolcu[isim] = numara
	}
	var takımadı string
	fmt.Println("Takımınız ismini giriniz: ")
	fmt.Scanln(&takımadı)
	fmt.Printf("%s Kamp Kadrosu--", takımadı)
	fmt.Println(&futbolcu)
}
