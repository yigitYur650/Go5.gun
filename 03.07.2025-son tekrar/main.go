/* package main

import (
	"fmt"
)

func main() {
	kanal := make(chan string)
	kanal1 := make(chan string)
	kanal2 := make(chan string)
	go cagri(kanal)
	fmt.Println(<-kanal)
	go temsilci1(kanal1)
	go temsilci2(kanal2)
	secim(kanal1, kanal2)
}
func cagri(kanal chan string) {
	kanal <- "Çağrı merkezine hoşgeldiniz yapmak istediğiniz işlemi seçiniz"
}
func secim(kanal1, kanal2 chan string) {
	var islem int
	fmt.Println("1-Fatura bedeli\n2-Teknik destek\n3-Kayıp, çalıntı vb durumlar\n4-Müşteri temsilcisine bağlan")
	fmt.Scanln(&islem)
	if islem == 1 {
		fmt.Println("Fatura bedeliniz tarifenizi satın alırken imzaladığınız sözleşmeye +25 TL eklenmiş fiyattır")
	}
	if islem == 2 {
		fmt.Println("Teknik servis şuan kullanım dışıdır")
	}
	if islem == 3 {
		fmt.Println("Geçmiş olsun")
	}
	if islem == 4 {
		select {
		case msg := <-kanal1:
			fmt.Println(msg)
		case mgt := <-kanal2:
			fmt.Println(mgt)
		}

	}
}
func temsilci1(kanal1 chan string) {
	kanal1 <- "Merhaba ismim Umutcan size nasıl yardımcı olabilirim"
}
func temsilci2(kanal2 chan string) {
	kanal2 <- "Merhaba ismim Güven size nasıl yardımcı olabilirim"
}
*/ /* package main

import (
	"fmt"
)

func main() {
	kanal := make(chan string)
	kanal1 := make(chan string, 1)
	kanal2 := make(chan string, 1)
	go cagri(kanal)
	fmt.Println(<-kanal)
	secim(kanal1, kanal2)
}
func cagri(kanal chan string) {
	kanal <- "Çağrı merkezine hoşgeldiniz yapmak istediğiniz işlemi seçiniz"
}
func secim(kanal1, kanal2 chan string) {
	var islem int
	fmt.Println("1-Fatura bedeli\n2-Teknik destek\n3-Kayıp, çalıntı vb durumlar\n4-Müşteri temsilcisine bağlan")
	fmt.Scanln(&islem)
	if islem == 1 {
		fmt.Println("Fatura bedeliniz tarifenizi satın alırken imzaladığınız sözleşmeye +25 TL eklenmiş fiyattır")
	}
	if islem == 2 {
		fmt.Println("Teknik servis şuan kullanım dışıdır")
	}
	if islem == 3 {
		fmt.Println("Geçmiş olsun")
	}
	if islem == 4 {
		go temsilci1(kanal1)
		go temsilci2(kanal2)
		for i := 0; i < 2; i++ {
			select {
			case msg := <-kanal1:
				fmt.Println(msg)
			case mgt := <-kanal2:
				fmt.Println(mgt)
			}
		}
	}
}
func temsilci1(kanal1 chan string) {
	kanal1 <- "Merhaba ismim Umutcan size nasıl yardımcı olabilirim"
}
func temsilci2(kanal2 chan string) {
	kanal2 <- "Merhaba ismim Güven size nasıl yardımcı olabilirim"
}
*/
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	kanal := make(chan string)
	kanal1 := make(chan string, 1)
	kanal2 := make(chan string, 1)
	go cagri(kanal)
	fmt.Println(<-kanal)
	secim(kanal1, kanal2)
}
func cagri(kanal chan string) {
	kanal <- "Çağrı merkezine hoşgeldiniz yapmak istediğiniz işlemi seçiniz"
}
func secim(kanal1, kanal2 chan string) {
	var islem int
	fmt.Println("1-Fatura bedeli\n2-Teknik destek\n3-Kayıp, çalıntı vb durumlar\n4-Müşteri temsilcisine bağlan")
	fmt.Scanln(&islem)
	if islem == 1 {
		fmt.Println("Fatura bedeliniz tarifenizi satın alırken imzaladığınız sözleşmeye +25 TL eklenmiş fiyattır")
	}
	if islem == 2 {
		fmt.Println("Teknik servis şuan kullanım dışıdır")
	}
	if islem == 3 {
		fmt.Println("Geçmiş olsun")
	}
	if islem == 4 {
		wg.Add(2)
		go temsilci1(kanal1)
		go temsilci2(kanal2)
		for i := 0; i < 2; i++ {
			select {
			case msg := <-kanal1:
				fmt.Println(msg)
			case mgt := <-kanal2:
				fmt.Println(mgt)
			}
		}
		wg.Wait()
	}
}
func temsilci1(kanal1 chan string) {
	defer wg.Done()
	kanal1 <- "Merhaba ismim Umutcan size nasıl yardımcı olabilirim"
}
func temsilci2(kanal2 chan string) {
	defer wg.Done()
	kanal2 <- "Merhaba ismim Güven size nasıl yardımcı olabilirim"
}
