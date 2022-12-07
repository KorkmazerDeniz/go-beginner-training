package main

import (
	"fmt"
	"math"
)

const PI = 3.14

func hesaplama(a, b int) (carp int, bol int) {
	carp = a * b
	bol = a / b
	return
}

func daire(yaricap float64) (alan float64, cevre float64) {

	alan = PI * (math.Pow(yaricap, 2.0))
	cevre = PI * 2 * yaricap
	return

}

func main() {

	/*

		carpim,bolme:=hesaplama(66,6)

		fmt.Println("66 X 6 =",carpim)
		fmt.Println("66 / 6 =",bolme)

	*/

	dairealan, dairecevre := daire(5.0)

	fmt.Println("Dairenin Alanı:", dairealan)
	fmt.Println("Dairenin Çevresi:", dairecevre)

}
