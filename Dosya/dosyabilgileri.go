package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	dosyaBilgi, err := os.Stat("veri.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Dosya Adı:", dosyaBilgi.Name())
	fmt.Println("Dosya İzinleri:", dosyaBilgi.Mode())
	fmt.Println("Dosya Boyutu:", dosyaBilgi.Size())
	fmt.Println("Değiştirme Tarihi:", dosyaBilgi.ModTime())
	fmt.Println("Klasör Mü?", dosyaBilgi.IsDir())
	fmt.Println("Sistem Interface", dosyaBilgi.Sys())

	if dosyaBilgi.Size() < 90 {
		os.Remove("veri.txt")
	}

}
