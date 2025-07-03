/*
	package main

import "fmt"

	func yemekSiparis(kanal chan string) {
		kanal <- "Yemeğiniz geldi\n"
	}

	func main() {
		//go routine kullanılan fonksiyonlarda return kullanılamaz
		kanal := make(chan string)
		go yemekSiparis(kanal)
		yemek := <-kanal
		fmt.Println("Yemeğiniz yola çıktı", yemek)
	}
*/
/* package main
import (
	"fmt"
	"time"
)
func mesaj(kanal chan string) {
	kanal <- "Kahveniz hazırlandı"
}
func main() {
	kanal := make(chan string, 1)
	fmt.Println("Kahve siparişiniz alındı")
	time.Sleep(2 * time.Second)
	go mesaj(kanal)
} */
/* package main

import (
	"fmt"
	"time"
)

func teknikServis(kanal chan string) {
	kanal <- "Müşteri temsilcisine bağlandınız"
}
func uyarı(kanal1 chan string) {
	kanal1 <- "Görüşmelerimiz kalite standartları gereği kayıt altına alınmaktadır"
}
func main() {
	kanal1 := make(chan string)
	kanal := make(chan string)
	fmt.Println("Teknik Servise hoşgeldiniz sizi müşteri temsilcisine bağlayacağız")
	time.Sleep(time.Second * 2)
	go uyarı(kanal1)
	fmt.Println(<-kanal1)
	time.Sleep(time.Second * 2)
	go teknikServis(kanal)
	fmt.Println(<-kanal)
} */
package main

import (
	"fmt"
)

func teknikServis(kanal chan string) {
	kanal <- "Müşteri temsilcisine bağlandınız"
}
func uyarı(kanal1 chan string) {
	kanal1 <- "Görüşmelerimiz kalite standartları gereği kayıt altına alınmaktadır"
}
func main() {
	kanal1 := make(chan string)
	kanal := make(chan string)
	fmt.Println("Teknik Servise hoşgeldiniz sizi müşteri temsilcisine bağlayacağız")
	go teknikServis(kanal)
	go uyarı(kanal1)
	for i := 0; i < 2; i++ {
		select {
		case mst := <-kanal1:
			fmt.Println(mst)
		case mst1 := <-kanal:
			fmt.Println(mst1)
		}
	}
}
