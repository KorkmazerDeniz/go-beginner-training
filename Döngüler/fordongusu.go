package main

import "fmt"

func main() {
	//x := 0

	//for x < 5 {
	//fmt.Println(x)
	//x += 1
	//}

	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("Çift Sayılar:%d\n", i)
		}
	}
}
