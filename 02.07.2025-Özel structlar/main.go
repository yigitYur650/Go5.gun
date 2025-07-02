package main

import (
	"fmt"
)

type film interface {
	Cokkötü(kötü string) string
}

type film1 interface {
	Cokiyi(iyi string) string
}

type harrypotter struct{}
type karaipkorsanları struct{}
type marvel struct{}
type dc struct{}

func (harrypotter) Cokkötü(kötü string) string {
	return fmt.Sprint("Bu film çok kötü")
}

func (karaipkorsanları) Cokkötü(kötü string) string {
	return fmt.Sprint("Bu film kötü")
}

func (marvel) Cokiyi(iyi string) string {
	return fmt.Sprint("Bu film çok iyi")
}

func (dc) Cokiyi(iyi string) string {
	return fmt.Sprint("Bu film iyi")
}

func main() {
	var h film
	var y film1

	h = harrypotter{}
	y = marvel{}

	// film interface için type switch
	switch f := h.(type) {
	case harrypotter:
		fmt.Println("Harry Potter türü film: ", f.Cokkötü(""))
	case karaipkorsanları:
		fmt.Println("Karaip Korsanları türü film: ", f.Cokkötü(""))
	default:
		fmt.Println("Bilinmeyen film türü")
	}

	// film1 interface için type switch
	switch f := y.(type) {
	case marvel:
		fmt.Println("Marvel türü film: ", f.Cokiyi(""))
	case dc:
		fmt.Println("DC türü film: ", f.Cokiyi(""))
	default:
		fmt.Println("Bilinmeyen film türü")
	}
}
