package main

import (
	"fmt"
	"time"
)

func main() {

	//2 adet bool değer taşıyan bir kanal oluşturalım
	k := make(chan bool, 2)

	//asenkron bir iş parçacığı oluşturalım
	go func() {

		//5 sn beklesin
		time.Sleep(time.Second * 5)

		//k kanalına bool bir değer gönderelim
		k <- true

		//tekrardan 2 sn beklesin
		time.Sleep(time.Second * 2)

		//ve k kanalına 2. değer de gönderilsin.
		k <- false
	}()

	//ana iş parçacığı k kanalına 2 değer gelene kadar bekleyecek
	fmt.Println(<-k, <-k) //çıktı: true false
	//iki bool değeri de baştırmak için k kanalını 2 defa yazdık
}
