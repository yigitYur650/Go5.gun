package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go ve()
	wg.Wait()
	fmt.Println("Williy Wonka")

}
func ve() {
	defer wg.Done()
	fmt.Println("Selam sana dost gün sana merhaba diyor")
}
