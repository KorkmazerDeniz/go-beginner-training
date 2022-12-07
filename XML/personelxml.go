package main

import (
	"encoding/xml"
	"io"
	"log"
	"os"
)

type Sirket struct {
	XMLName     xml.Name   `xml:"sirket"`
	Personeller []Personel `xml:"personel"`
}

type Personel struct {
	XMLName  xml.Name `xml:"personel"`
	TCNO     int      `xml:"tcno"`
	PerAd    string   `xml:"personelAd"`
	PerSoyad string   `xml:"personelSoyad"`
	PerMevki string   `xml:"mevki"`
}

func main() {

	dosya, err := os.Create("personel.xml")
	if err != nil {
		log.Fatal(err)
	}

	str := "<?xml version=\"1.0\" encoding=\"utf-8\"?>\n"
	dosya.WriteString(str)

	s := &Sirket{}

	s.Personeller = append(s.Personeller, Personel{TCNO: 23296, PerAd: "Deniz", PerSoyad: "Korkmazer", PerMevki: "Yazılımcı"})
	s.Personeller = append(s.Personeller, Personel{TCNO: 38954, PerAd: "Ali", PerSoyad: "Özerinç", PerMevki: "Yönetici"})

	xmlWriter := io.Writer(dosya)

	enc := xml.NewEncoder(xmlWriter)
	enc.Indent("  ", "    ") //Girinti oluşturma
	enc.Encode(s)

	//XML TO JSON

	//jsonVeri, err := json.Marshal(s)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//jsonDosya, err := os.Create("yeni.json")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//defer jsonDosya.Close()
	//
	//jsonDosya.Write(jsonVeri)

}
