package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type kitap struct {
	ad    string
	yazar string
	yıl   int
	konu  string
}

func main() {
	var n int
	var yıl int
	kütüp := []kitap{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Kaç kitap kayıt edeceksiniz: ")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Println("Kitabın ismini giriniz")
		ad, _ := reader.ReadString('\n')
		ad = strings.TrimSpace(ad)
		fmt.Println("Kitabın yazarının ismini giriniz")
		yazar, _ := reader.ReadString('\n')
		yazar = strings.TrimSpace(yazar)
		fmt.Println("Kitabın konusunu giriniz\n(Korku,komedi,psikoloji)")
		konu, _ := reader.ReadString('\n')
		konu = strings.TrimSpace(konu)
		fmt.Println("Kitabın kaç yılında basıldığını giriniz: ")
		fmt.Scanf("%d\n", &yıl)
		yenikitap := kitap{ad: ad, yazar: yazar, konu: konu, yıl: yıl}
		kütüp = append(kütüp, yenikitap)

	}
	fmt.Println(kütüp)
}
