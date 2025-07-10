package main

func main() {
	// x := 10
	// y := 17.8
	// fmt.Printf("%v %T\n", x, x)
	// fmt.Printf("%v %T\n", y, y)
	// fmt.Println(x + int(y)) // satır bazlı bir tip dönüşümü var
	// fmt.Printf("%v %T\n", y, y)
	// var x int8 = 10
	// var y int16 = 10
	// fmt.Println(x + y) // tipleri değişik olduğu için aynı sayıda olsa toplanmıyor
	// var x int16 = 130
	// var y int8
	// y = int8(x)
	// println(y)
	// println(x)
	num1 := 107
	str1 := string(num1) // int türünden string türüne dönüştürme
	println(str1)
}

// int(y) şeklinde yazdığımda y int türüne dönüştü ve ondalıklı kısmı 0 a yuvarlandı
// int flaot a mı dönmeli yoksa tam tersi mi olmalı sorusunun cevabı kodun o anki gerekliliğine göre değişiklik gösterir
