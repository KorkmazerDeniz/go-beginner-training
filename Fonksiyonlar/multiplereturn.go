package main

import "fmt"

/* snytax
fun fonksiyon_adı(parametre1,parametre2,...)(dönüş veri tipi1,dönüş veri tipi2){
// işlemler..
*/

func ucislem(x, y int) (int, int, int) {
	return x - y, x + y, x * y

}
func maxmin(a, b int) (int, int) {
	if a > b {
		return a, b
	} else {
		return b, a
	}
}

func main() {
	//sayi1, sayi2, sayi3 := ucislem(5, 9)
	//
	//fmt.Printf("Sonuc:%d \n", sayi1)
	//fmt.Printf("Sonuc:%d \n", sayi2)
	//fmt.Printf("Sonuc:%d \n", sayi3)

	max, min := maxmin(30, 60)

	fmt.Printf("Max=%d Min=%d", max, min)
}
