package mainpackage main

import "fmt"

func start(){
	defer fmt.Println("PC AÇILDI")
	fmt.Println("PC AÇILIYOR...")
}

func stop(){
	defer fmt.Println("PC KAPANDI.")
	fmt.Println("PC KAPANIYOR...")


}

func sira(){

	defer fmt.Println("Birinci")
	defer fmt.Println("İkinci")
	defer fmt.Println("Üçüncü")

	fmt.Println("Fonksiyon Başlıyor...")
}



func main(){


	/*
		defer fmt.Println("Merhaba")
		fmt.Println("Dünya")


	*/
	/*
		start()
		stop()

	*/
	sira()











}

