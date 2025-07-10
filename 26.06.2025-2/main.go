package main

import (
	"fmt"
)

func main() {
	const aşk string = "Golang"
	fmt.Println("Bugün", aşk, "dilini öğreniyorum.") // <3 */
	const pi float64 = 3.14
	r := 7.0
	fmt.Println("Çemberin alanı ", pi*(r*r)) // burada bir değerin float bir değerin int olmaması önemli yoksa derlemede hata alırız
	fmt.Println("Çemberin çevresi ", 2*pi*r)
	
}
