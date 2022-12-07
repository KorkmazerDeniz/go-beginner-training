package main

import "fmt"

func main() {

	//yeni bir oluşturma

	var a []int = []int{11, 22, 33, 44, 55}

	//fmt.Printf("%T",a)

	fmt.Println(a)

	//append

	/*
		b:=append(a,66)
		fmt.Println(b)

	*/

	a = append(a, 66)

	fmt.Println(a)

	a = append(a, 77)

	fmt.Println(a)

	var sehirler = []string{"Malatya", "Ankara", "Bursa"}
	sehirler = append(sehirler, "Antalya")
	sehirler = append(sehirler, "İzmir")

	fmt.Println(sehirler)

	//make

	sayilar := make([]int, 5)

	fmt.Println(sayilar)

	sayilar[0] = 3
	sayilar[1] = 6
	sayilar[2] = 9
	sayilar[3] = 12
	sayilar[4] = 15

	fmt.Println(sayilar)

	sayilar = append(sayilar, 18)

	fmt.Println(sayilar)

	fmt.Println("Kapasite:", cap(sayilar))
	fmt.Println("Uzunluk:", len(sayilar))

}
