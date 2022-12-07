package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("1-Toplama\n" +
		"2-Çıkarma\n" +
		"3-Çarpma\n" +
		"4-Bölme\n" +
		"Lütfen İşlem Seçiniz")
	scanner.Scan()

	secim := scanner.Text()

	fmt.Print("1.Sayıyı Giriniz")
	scanner.Scan()
	sayi1, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("2.Sayıyı Giriniz:")
	scanner.Scan()
	sayi2, _ := strconv.ParseFloat(scanner.Text(), 64)

	switch secim {
	case "1":
		fmt.Println("Toplama:", sayi1+sayi2)
	case "2":
		fmt.Println("Çıkarma:", sayi1-sayi2)
	case "3":
		fmt.Println("Çarpma:", sayi1*sayi2)
	case "4":
		fmt.Println("Bölme:", sayi1/sayi2)
	default:

		fmt.Println("Hatalı Seçim Yaptınız")

	}

}
