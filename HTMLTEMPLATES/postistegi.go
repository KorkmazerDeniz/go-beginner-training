package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	address := "https://orneksite.com/kisiler"
	hamVeri := map[string]string{
		"isim":    "Deniz",
		"soyisim": "Korkmazer",
	}
	jsonVeri, hata := json.Marshal(hamVeri)
	//jsonVeri string tipine çevirildiğinde:
	//{"isim":"Kaan","soyisim":"Kuşcu"}

	if hata != nil {
		panic(hata)
	}

	//Aşağıdaki gönderdiğimiz verinin json tipinde olduğunu bildiriyoruz.
	sorgu, hata := http.Post(address, "application/json", bytes.NewBuffer(jsonVeri))
	if hata != nil {
		panic(hata)
	}

	//sonucu sonuç değişkenine atayalım
	sonuç, hata := io.ReadAll(sorgu.Body)
	if hata != nil {
		panic(hata)
	}

	//Son olarak ekrana bastıralım.
	fmt.Println(string(sonuç))
}
