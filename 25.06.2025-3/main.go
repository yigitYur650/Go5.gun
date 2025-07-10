package main

import "fmt"

func main() {
	//print-println-printf arasındaki farklar:
	//1. print ve println, verilen argümanı ham şekilde yazdırırken,
	//   printf formatlama yaparak yazdırır
	name := "John"
	fmt.Print("Hello, ", name, "!")
	fmt.Println(" ")
	fmt.Println("Hello,", name, "!") // alt satıra atıyor
	fmt.Printf("Hello, ", name, "!") // alt satıra atmaz
}
