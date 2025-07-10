package main

import "fmt"

func main() {
	fmt.Println("Akdeniz Üniversitesi not hesaplama robotuna hoşgeldiniz.\nOkulumuzda not dağılımı %25 Proje, %25 Vize, %50 Final şeklindedir.")
	var not1 float64 // bir not burda := şekilde tanımladığımda hep hata aldım kısayolunu hala öğrenemedim
	var not2 float64
	var proje1 float64
	fmt.Println("Vize notunuzu giriniz:  ")
	fmt.Scanln(&not1)
	fmt.Println("Final notunuzu giriniz:  ")
	fmt.Scanln(&not2)
	fmt.Println("Proje notunuzu giriniz:  ")
	fmt.Scanln(&proje1)
	ort := ortalama(not1, not2, proje1)
	harfDagilimi(ort)
	passed(ort)
}
func ortalama(not1 float64, not2 float64, proje1 float64) float64 {
	ort := (not1 * 0.25) + (proje1 * 0.25) + (not2 * 0.5)
	//fmt.Println("Ortalama:  ", ort)
	return ort
}
func harfDagilimi(ortalama float64) {
	switch true {
	case ortalama >= 90:
		fmt.Println("AA")
	case ortalama >= 85:
		fmt.Println("BA")
	case ortalama >= 80:
		fmt.Println("BB")
	case ortalama >= 75:
		fmt.Println("CB")
	case ortalama >= 70:
		fmt.Println("CC")
	case ortalama >= 65:
		fmt.Println("DC")
	case ortalama >= 60:
		fmt.Println("DD")
	case ortalama >= 50:
		fmt.Println("FD")
	default:
		fmt.Println("FF")
	}
}
func passed(ortalama float64) {
	if ortalama < 65 {
		fmt.Println("KALDI")
	} else {
		fmt.Println("GEÇTİ")
	}
}
