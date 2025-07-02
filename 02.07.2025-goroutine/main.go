package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // Burda waitGroupu tanımladık çünkü;
//Eğer tanımlamasaydık x ve y yazımı rastgele bir biçimde olacaktı

func main() {
	wg.Add(1) // Burda beklenmesi gereken bir değer olduğunu söyledik
	go Printx()
	wg.Wait() // burda üstteki go bitmeden aşağı geçmemesini söyledik
	Printy()
}
func Printx() {
	for i := 0; i < 20; i++ {
		fmt.Println("Y")
	}
	wg.Done() // Burdada döngünün bittiği sinyalini aldık
}
func Printy() {
	for i := 0; i < 20; i++ {
		fmt.Println("X")
	}
}
