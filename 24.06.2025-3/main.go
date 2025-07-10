package main

import "fmt"

func main() {
	//var name string = "Rasul"
	var age int = 13
	var okul bool = true

	var name = "Rasul" // Veri tipi berlirtmedende yazabiliyoruz
	var weigght = 65.5 // float64 float32 de olabilir
	fmt.Println(name, age, okul, weigght)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", okul)
	fmt.Printf("%T\n", weigght)

}
