package main

import (
	"fmt"
)

type telefon struct {
	model     string
	cikisyili int
	marka     string
}

func (t telefon) teknosa() {
	fmt.Printf("Telefonun markası:%s\n Modeli:%s\n Çıkış yılı:%d", t.marka, t.model, t.cikisyili)
}
func main() {
	var t telefon
	fmt.Println("Telefonunuzun markasını giriniz: ")
	fmt.Scanln(&t.marka)
	fmt.Println("Telefonunuzun modelini giriniz: ")
	fmt.Scanln(&t.model)
	fmt.Println("Telefonunuzun çıkış yılını giriniz: ")
	fmt.Scanln(&t.cikisyili)
	t.teknosa()
}
