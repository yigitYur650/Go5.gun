package main

import "fmt"

var packVar = "Package Scope"

func main() {
	var funcVar = "Func Scope"
	fmt.Println(funcVar)
	fmt.Println(packVar)
	// Bu derste scope yani anladığım kadarıyla package dosyası package içinde olacak func ise funcun içinde işte if blokları varsa onlarda kendi içlerinde olacak
}
