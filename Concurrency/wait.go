package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine")
		}

		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Ana Thread")

}
