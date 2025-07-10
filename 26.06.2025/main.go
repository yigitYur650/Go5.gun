// 1 : int x, float64 y type conversion sample

// package main

// import "fmt"

// func main() {
// 	var y float64 = 3.14
// 	var x int
// 	y = float64(x)

// 	fmt.Println(x)
// 	fmt.Println(y)
// 	fmt.Printf("%T\n %T", x, y) // Burada normalde iki değeride 0 çıkartan bir kod olması gerekiyordu ancak o kod rahatlıkla yazılabildiği için ben integer x değerini float64' çevirerek yaptım.
// }

// 2 : multiple assing sample x, y = y, x
// package main

// import "fmt"

// func main() {
// 	x := 2
// 	y := 7
// 	fmt.Println("X:", x, "Y:", y) // x ve y yi tanımladım sonrada onların değerlerini değiştirdim.
// 	x, y = y, x
// 	fmt.Println("X:", x, "Y:", y)
// }

// 3 : non English variable names

// go dilinde değişken isimleri İngilizce dışında da kullanılabilir. mesela çince falan

// 4 : shadowing kavramı? gölgeleme
// package main

// import "fmt"

// func main() {
// 	x := 7
// 	if true {
// 		x := 2               // burada x değişkenini sadece = la tanımlasaydım en alttaki printlnde de 2 çıkacaktı
// 		fmt.Println("X:", x) // burada x değişkenini bi if bloğunda bide normalde tanımladım ifteki ayrı öbürü ayrı çalıştı
// 	}
// 	fmt.Println("X:", x)
//}

// 5 : 40 as a string
// package main

// import "fmt"

// func main() {
// 	x := 66
// 	yiit := string(x)
// 	fmt.Printf("x: %d, yiit: %s\n", x, yiit) // burada da integer değeri ASCII koda göre stringe çevirdim
// 	fmt.Printf("x: %T, yiit: %T\n", x, yiit)
// }
