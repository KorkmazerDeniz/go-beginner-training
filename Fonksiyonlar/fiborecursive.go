package main

import "fmt"

//Fibonnacci

//0,1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, …
//a b  a  b
//f(n)=f(n-1)+f(n-2)

func fibo(n int) int {
	a := 0
	b := 1

	for i := 0; i < n; i++ {
		gecici := a
		a = b
		b = gecici + a
	}
	return a

}

func main() {

	for n := 0; n <= 20; n++ {
		sonuc := fibo(n)
		fmt.Printf("Fibonacci:%d =%d \n", n, sonuc)
	}

}
