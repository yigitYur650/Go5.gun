package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type film struct {
	filmAdi  string
	filmKonu string
	yonetmen string
	cikisyil int
}

func (f film) sinema() {
	fmt.Printf("Filmin ismi:%s\nFilmin konusu:%s\nFilmin yönetmeni:%s\nFilmin Çıkış yılı:%d\n", f.filmAdi, f.filmKonu, f.yonetmen, f.cikisyil) // filmleri yazdırmak için struct maini kullandım
}
func main() {
	var f film
	var filmler []film
	reader := bufio.NewReader(os.Stdin) // burda bufio ile girilen deeğerlerin arasında boşluk bırakabilmeyi sağladım
	fmt.Printf("Çıkmak için q yazınız. Filmin çıkış yılı girilirken geçerli değildir\n")

	for {
		fmt.Println("Bir film ismi giriniz: ")
		f.filmAdi, _ = reader.ReadString('\n')
		f.filmAdi = strings.TrimSpace(f.filmAdi)
		if f.filmAdi == "q" { // burda çıkmak için komut bıraktım
			break
		}
		fmt.Println("Filmin konusunu giriniz: ")
		f.filmKonu, _ = reader.ReadString('\n')
		f.filmKonu = strings.TrimSpace(f.filmKonu)
		if f.filmKonu == "q" {
			break
		}
		fmt.Println("Filmin yönetmeninin ismini giriniz: ")
		f.yonetmen, _ = reader.ReadString('\n')
		f.yonetmen = strings.TrimSpace(f.yonetmen)
		if f.yonetmen == "q" {
			break
		}
		fmt.Println("Filmin çıkış tarihini giriniz: ")
		_, err := fmt.Scanln(&f.cikisyil) //Burada hata kontrol çalıştırdım
		if err != nil {
			fmt.Println("Lütfen geçerli bir sayı giriniz")
			continue
		}
		filmler = append(filmler, f) //burda filmleri filmler esnek dizisine aldım ancak buradan yazdırmadım o zaman kodum karıır veya struct ile yaptığım fonksiyonu silmem gerekirdi
		f.sinema()
	}
}
