package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	kaynak, err := ioutil.ReadFile("veri.txt")
	if err != nil {
		log.Fatal(err)
	}

	yeniDosya, err2 := os.Create("yenidosya.txt")
	if err2 != nil {
		log.Fatal(err2)
	}

	err3 := ioutil.WriteFile("yenidosya.txt", kaynak, 0666)
	if err3 != nil {
		fmt.Println("Oluştururken hata oluştu", err3)
	}

	defer yeniDosya.Close()

}
