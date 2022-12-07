package main

import "fmt"

func recursive() {
	fmt.Println("Merhaba")
	recursive()
}

func faktoriyel(sayi int) int {
	if sayi <= 1 {
		return 1
	}

	return sayi * faktoriyel(sayi-1)

}

func main() {

	/*
		recursive()


	*/

	sayi := 12 // 12x11x10....1

	fmt.Printf("%d sayısının Faktoriyeli: %d", sayi, faktoriyel(sayi))

}
