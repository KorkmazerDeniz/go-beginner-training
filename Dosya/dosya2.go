package main

import (
	"log"
	"os"
)

func main() {

	//eskiAd := "veri.txt"
	//yeniAd := "data.txt"
	//
	//os.Rename(eskiAd, yeniAd)

	dosyaKonum := "data.txt"
	yeniKonum := "Trym/veri.txt"

	err := os.Rename(dosyaKonum, yeniKonum)
	if err != nil {
		log.Fatal(err)
	}

}
