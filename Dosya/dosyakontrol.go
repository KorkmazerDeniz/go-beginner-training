package main

import (
	"fmt"
	"os"
)

func main() {

	_, err := os.Stat("yeni.txt")
	if err != nil {
		fmt.Println("HATA:", err)
	}

	dosyaKontrol := os.IsNotExist(err)
	fmt.Println(dosyaKontrol)

	if dosyaKontrol == true {
		os.Create("veri.txt")
	}

}
