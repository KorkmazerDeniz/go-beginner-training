package main

import "fmt"

func main() {

	isim := "Deniz1"
	if isim == "Deniz" {
		fmt.Println("Tekrar Hoş Geldin Deniz")
	} else {
		fmt.Println("Lütfen çıkış yapınız")
	}

	yas := 23

	fmt.Println("18 Yaşından Küçükler Giremez")

	if yas >= 18 {
		fmt.Println("Mekana Hoşgeldiniz")
	} else if yas < 18 && yas > 19 {
		fmt.Println("Ailenle girebilirsin")
	} else {
		fmt.Println("GİREMEZSİNİZ")
	}
}
