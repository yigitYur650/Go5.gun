package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	kisiler := make(map[string]int)
	var kac int
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Kaç futbolcuya lisans çıkartmak istiyorsunuz: ")
	fmt.Scanln(&kac)
	for i := 0; i < kac; i++ {
		fmt.Printf("%d. Sporcunun adını ve soyadını giriniz :  ", i+1)
		adSoyad, _ := reader.ReadString('\n')
		adSoyad = strings.TrimSpace(adSoyad)
		fmt.Printf("%d. Sporcunun forma numarasını giriniz:  ", i+1)
		var numara int
		fmt.Scanf("%d\n", &numara)
		kisiler[adSoyad] = numara
	}
	var takim string
	fmt.Println("Takımınızın adını giriniz:  ")
	fmt.Scanln(&takim)
	fmt.Printf("******%s Kadro Listesi******\n", takim)
	fmt.Println(kisiler)
}
