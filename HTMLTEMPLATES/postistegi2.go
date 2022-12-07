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

	//struct içindeki değişken isimlerini büyük harfle
	//yazmazsak json'a çevirirken hata alırız.
	structVeri := struct {
		İsim    string `json:"isim"`    //json'da isim olacak
		Soyisim string `json:"soyisim"` //json'da soyisim olacak
		Yaş     int    `json:"yas"`     //json'da yas olacak
	}{"Deniz", "Korkmazer", 23}

	//struct'ımızı json'a çevirelim.
	jsonVeri, hata := json.Marshal(structVeri)
	//jsonVeri string tipine çevirildiğinde:
	//{"isim":"Deniz","soyisim":"Korkmazer","yas":23}

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
