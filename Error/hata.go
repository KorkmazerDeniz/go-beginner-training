package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	2 paket vardır 1-errors 2-fmt
	nill => Null

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Sayı Giriniz")

		a, err := strconv.ParseInt(scanner.Text(), 10, 64)

		if err != nil {
			//fmt.Println(err.Error())
			log.Fatal("HATA", err)
		}

		fmt.Println(a)
	}
	file, err := os.Open("/main.go")
	if err != nil {
		fmt.Println(errors.New("DOSYA YOK"))
	}
	fmt.Println(file.Name(), "BAŞARILI BİR ŞEKİLDE AÇILDI")

}
