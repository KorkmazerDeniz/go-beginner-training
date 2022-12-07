package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Students struct {
	XMLNAME        xml.Name `xml:"Ogrenciler"`
	OgrenciListesi []Ogr    `xml:"ogrenci"`
}
type Ogr struct {
	//ALT 96 YADA CTRL VE ALT , TUŞUNA BASILCAK
	XMLNAME      xml.Name `xml:"ogrenci"`
	OgrenciNo    int      `xml:"ogrNo"`
	OgrenciAd    string   `xml:"ogrAd"`
	OgrenciSoyad string   `xml:"ogrSoyad"`
	OgrenciNot   int      `xml:"not"`
}

func main() {

	dosya, err := os.Open("ogrenciler.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer dosya.Close()

	arr, err := ioutil.ReadAll(dosya)
	if err != nil {
		log.Fatal(err)
	}

	//var ogrenci Ogr
	//xml.Unmarshal(arr, ogrenci)
	//fmt.Println(string(arr))

	Ogrenci := &Students{}
	xml.Unmarshal(arr, Ogrenci)

	for _, degerler := range Ogrenci.OgrenciListesi {
		fmt.Println("Öğrenci Ad:", degerler.OgrenciAd)
		fmt.Println("Öğrenci Not:", degerler.OgrenciNot)
	}

}
