package main

import "fmt"

//built-in func:gömülü fonk.
//user-defined func

func merhaba() { //Parametresiz Func.
	fmt.Println("Merhaba")
}

func merhabaPar(mesaj string) { //Parametreli Func
	fmt.Println("Merhaba ", mesaj)
}

func main() {

	merhaba()
	merhaba()
	merhaba()

	merhabaPar("Ferhat")
	merhabaPar("Deniz")
}
