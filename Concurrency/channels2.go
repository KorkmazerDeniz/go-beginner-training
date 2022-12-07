package main

import (
	"time"
)

func main() {

	//bir kanal oluşturalım
	k := make(chan bool)
	//bu kanalımız bool değer taşıyacak

	//asenkron bir iş parçacığı oluşturalım
	go func() {

		//bu iş parçacığı 5 sn beklesin
		time.Sleep(time.Second * 5)

		//k kanalına bool bir değer gönderelim
		k <- true
	}()

	//ana iş parçacığı k kanalına değer gelene kadar bekleyecek
	<-k
	//değer geldiğinde program sonlanacaktır.
}
