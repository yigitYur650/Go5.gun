package main

import "fmt"

func main() {
	/* x := 5.00
	y := 2.00
	fmt.Println("X+Y:", x+y, "X-Y:", x-y) // bouble yazıp ne olacağına bakacam
	//fmt.Println("X*Y:", x*y, "X/Y:", x/y)
	fmt.Printf("X/Y: %v %T\n", x/y, x/y)
	fmt.Printf("X*Y: %v %T", x*y, x*y) // burda ondalıklı yazdığım değerlerin bölmede veri tipinin farklı çarpmada farklı olduğunu gördüm (bunu kendim denedim videoda yoktu ilerleyen kısımlarında varmıs ama bu benim bu çözümü kendim bulduğum gerçeğini değiştirmez)
	*/
	x := 5
	fmt.Println(x)
	x = x + 1
	fmt.Println(3 * x)
	x++ // operatör hep değişkenden sonra gelir
	fmt.Println(x)
	x--
	fmt.Println(x)
	//fmt.Println(x++)// böyle bir kullanım yoktur yukardada yazıdğım gibi -- yada ++ operatörleri değişkenden sonra gelir
}
