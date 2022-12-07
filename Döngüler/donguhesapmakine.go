package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	for true {

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("\nİŞLEM SEÇİNİZ\n" +
			"1-Toplama\n" +
			"2-Çıkarma\n" +
			"3-Çarpma\n" +
			"4-Bölme\n" +
			"ÇIKIŞ İÇİN q'ya basınız..")

		scanner.Scan()
		secim := scanner.Text()

		if secim == "q" || secim == "Q" {
			fmt.Println("Çıkış yapılıyor...")
			break
		}

		fmt.Print("1. sayıyı Giriniz:")
		scanner.Scan()
		sayi1, _ := strconv.ParseFloat(scanner.Text(), 64)

		fmt.Print("2. sayıyı Giriniz:")
		scanner.Scan()
		sayi2, _ := strconv.ParseFloat(scanner.Text(), 64)

		if secim == "1" {
			fmt.Printf("%.2f + %.2f = %.2f", sayi1, sayi2, sayi1+sayi2)
		} else if secim == "2" {
			fmt.Printf("%.2f - %.2f = %.2f", sayi1, sayi2, sayi1-sayi2)
		} else if secim == "3" {
			fmt.Printf("%.2f X %.2f = %.2f", sayi1, sayi2, sayi1*sayi2)
		} else if secim == "4" {
			fmt.Printf("%.2f / %.2f = %.2f", sayi1, sayi2, sayi1/sayi2)
		} else {
			fmt.Println("Hatalı Giriş")
		}

	}

}
