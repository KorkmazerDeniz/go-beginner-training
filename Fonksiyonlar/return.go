package main

import "fmt"

func toplama(x int, y int) int {
	return x + y
}

func onlacarp(t int) int {
	return t * 10
}

func bolme(a, b float32, c string) float32 {
	return a / b
}

func ikiylecarp(g int) int {
	return g * 2
}

func main() {

	fmt.Println(toplama(7, 8))

	a := onlacarp(8)
	fmt.Println(a)

	fmt.Println(bolme(22.7, 13.3, "AHKS"))

	for i := 0; i <= 10; i++ {
		fmt.Println(ikiylecarp(i))
	}

}
