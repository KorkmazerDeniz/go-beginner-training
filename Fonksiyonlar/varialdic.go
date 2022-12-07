package main

import "fmt"

func merhaba1(hello ...string) {
	fmt.Println(hello)
	fmt.Println(hello[2])
}

func sehirler(param ...string) {
	for _, sehir := range param {
		fmt.Println(sehir)
	}
}

func toplam(sayilar ...int) int {
	sonuc := 0
	for _, sayi := range sayilar {
		sonuc += sayi //sonuc=sonuc+sayi
	}

	return sonuc
}

func main() {

	/*
		merhaba1("Ahmet","Buse","Cem","Betül","Ali")
	*/

	/*
		sehirler("Çanakkale","İzmir","Bayburt","Muş")


	*/

	son := toplam(5, 6, 4, 3, 6, 7, 4, 3, 5, 6, 87, 98, 7)

	fmt.Println(son)

}
