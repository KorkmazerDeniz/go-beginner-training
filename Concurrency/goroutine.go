package main

import (
	"fmt"
	"time"
)

func yazdir(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {

	//Thread 1 MB  Goroutine 2KB

	go yazdir("Merhaba")
	time.Sleep(1 * time.Second)

	yazdir("Dunya")

}
