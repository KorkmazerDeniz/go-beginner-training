package main

import "fmt"

func main() {

	var isim string
	var yas int

	fmt.Print("Adınızı Giriniz:")
	fmt.Scan(&isim)

	fmt.Print("Yaşınızı Giriniz:")
	fmt.Scan(&yas)

	fmt.Printf("Adınız:%s Yaşınız:%d", isim, yas)

}
