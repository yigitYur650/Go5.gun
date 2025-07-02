package main

import "fmt"

type git interface {
	Git(nasıl string) string
}
type araba struct{}
type ucak struct{}
type gemi struct{}

func (a araba) Git(nasıl string) string {
	return fmt.Sprint("Araba tekerleklerinin üzerinde gider")
}
func (u ucak) Git(nasıl string) string {
	return fmt.Sprint("Uçak kanatları ile gider")
}
func (g gemi) Git(nasıl string) string {
	return fmt.Sprint("Gemi fizik kurallarına göre yüzeyi ile batmaz ve motorlarının güç verdiği pervanelerle gider")
}
func main() {
	var araç git
	araç = araba{}
	switch a := araç.(type) {
	case araba:
		fmt.Println(a.Git(""))
	case ucak:
		fmt.Println(a.Git(""))
	case gemi:
		fmt.Println(a.Git(""))
	}

}
