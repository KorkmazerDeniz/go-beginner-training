package main

import "fmt"

func main() {

	dizi := [5]int{1, 3, 5, 7, 9}

	var dilim []int = dizi[0:5]

	/*

		dizi[0:10]
		dizi[:10]

		dizi[0:]

		dizi[:]


	*/

	fmt.Println(dilim)

	a := [5]string{"Go", "Python", "Kotlin", "PHP", "JavaScript"}
	b := a[1:3]

	fmt.Println("Slice:", b)

	fmt.Println("Slice Kapasitesi:", cap(b))

	fmt.Println("Slice Uzunluğu:", len(b))

	fmt.Println("Array:", a)

}
