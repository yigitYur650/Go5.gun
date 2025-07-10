// 1-) studentName --> John Doe, grade --> 77, isPassed --> true değişkenlerini
// 3 farklı yöntem ile oluşturup, çıktısını yazdırınız.
/* package main

import "fmt"

func main() {
	name, grade, isPassed := "Feriha", 77, true //1- Değişkenleri tek satır içersinde tanımladım
	fmt.Println("Name:", name)
	fmt.Println("Grade:", grade)
	fmt.Println("Is Passed:", isPassed) */
//}
// package main

// import "fmt"

// func main() {
// var studentName string = "John Doe"
// var grade int = 77
// var isPassed bool = true

// fmt.Println("Student Name:", studentName) //2.Yöntemde uzun uzun tanımladım
// fmt.Println("Grade:", grade)
// fmt.Println("Is Passed:", isPassed)
// }
// package main

// import "fmt"

// func main() {

// 	var (
// 		name     = "John Doe" // 3.Yöntemde paranteze alarak bir tanımlaam yaptım
// 		grade    = 77
// 		isPassed = true
// 	)

// 	fmt.Println("Student Name:", name)
// 	fmt.Println("Grade:", grade)
// 	fmt.Println("Is Passed:", isPassed)

// 2-) yukarıda belirtilen değişkenleri tek satır içerisinde tanımlayınız.
// package main

// import "fmt"

// func main() {
// 	name, grade, isPassed := "John Doe", 77, true // Değişkenleri tek satırda tanımladım
// 	fmt.Println("Student Name:", name)
// 	fmt.Println("Grade:", grade)
// 	fmt.Println("Is Passed:", isPassed)
// }

// 3-) "Declaration", "Assign", "Initialization", "Initial Value" kavramlarını açıklayınız. (Terminoloji)
// package main

// import "fmt"

// func main() {
// var name string // Declaration: burda değişkeni deklare ettim ve kod çalışıyor zero values veriyor sadece
// fmt.Println(name)
// }
package main

func main() {
	var name string = "Yiğit" // Assign: Burada bir değişkene değer atadım assign de bu demek zaten
} // İnitialization: Değişkeni tanımlayıp, ona bir değer atamak onuda burda yaptım tekrarlamıyorum
// Initial Value: Değişkenin başlangıç değeri demek mesela name değişkenine altta "=" sembolle tekrar başka bir isim versek bu 2. değeri olur üsttekide initial value olur
// 4-) "Statically Typed" vs "Dynamically Typed" ifadelerini GO ve Python üzerinden gösteriniz.
// Burası için Python kodu yazmam gerekiyor onun yerine açıklamakla ve buraya çalışmayan bir örnek yazacağım
// Pyhtonda
//x=10 print(x)
//x="Selam" print(x) Python burda dinamik tipte olduu için iki değeride verecektir ancak bunu Go da yaparsak hata verecektir. Go statik tipli bir dildir

// 5-) ":=" vs "=" aradaki farkı gösteriniz, double declaration
// ":=" sembolü var değişkeni tanımlamadan değer atamamızı sağlar kısayol gibi he ama mesela sürekli geri döneceğimiz bir değer olursa var yine kullanılmalı öteki yandan
// "=" sembolü var olan bir değişkeni günceller yani az önceki python örneğinin aynı veri tipinden olduğunu varsayarsanız onu güncelliyor işte
