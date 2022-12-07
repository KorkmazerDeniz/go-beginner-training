package main

import "fmt"

func abc(t int) int {
	return t * 10
}

func main() {

	any := func(toplam int) int {
		return toplam / 3 * 4 * 5
	}(10)
	/* diğer fonksiyonlardaki gibi
	parametre verilmediği için fonksiyonun sonuna
	parametre verilir.*/

	fmt.Println(any)

	/*Bizim anonim fonksiyonları kullanma amacımız
	genellikle başka bir fonksiyonun içerisine argüman
	olarak göndermek içindir.

	*/

	fmt.Println(abc(any))

}
