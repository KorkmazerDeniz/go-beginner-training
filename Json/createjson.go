package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Companies struct {
	CListe []Sirket
}

type Sirket struct {
	ID            int
	Isim          string
	KurulusTarihi int
	DevamEdiyorMu bool
}

func main() {

	Sirket1 := Sirket{}
	Sirket1.ID = 1
	Sirket1.Isim = "Ermat"
	Sirket1.KurulusTarihi = 1989
	Sirket1.DevamEdiyorMu = true

	sirket1Byte, err := json.Marshal(Sirket1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(sirket1Byte))

	sirket2 := Sirket{55, "Mais", 1969, true}
	sirket2Byte, _ := json.Marshal(sirket2)

	fmt.Println(string(sirket2Byte))

	dosya, err := os.Create("sirketlerim.json")
	if err != nil {
		log.Fatal(err)
	}

	c1 := Companies{}

	c1.CListe = append(c1.CListe, Sirket{ID: 543, Isim: "Renault", KurulusTarihi: 1898, DevamEdiyorMu: false})
	c1.CListe = append(c1.CListe, Sirket{ID: 16, Isim: "Dacia", KurulusTarihi: 1966, DevamEdiyorMu: false})

	jsonWriter := io.Writer(dosya)
	enc := json.NewEncoder(jsonWriter)
	json.MarshalIndent(c1, " ", "\t")
	enc.Encode(c1)

}
