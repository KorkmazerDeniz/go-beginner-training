package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// yakalayıcı fonksiyonumuz
func anasayfa(w http.ResponseWriter, r *http.Request) {
	//isim değişkenimiz
	isim := "Kaan"

	//burada şablon oluşturuyoruz
	şablon, _ := template.ParseFiles("sablon.html")

	//Burada da şablonu çalıştırmasını ve isim
	//değişkenini kullanmasını istiyoruz.
	şablon.Execute(w, isim)
}
func main() {
	fmt.Println("Sunucu Başladı")

	//ana dizini anasayfa fonksiyonu ile yakalayalım
	http.HandleFunc("/", anasayfa)

	//portu 8000 yapalım ve sunucuyu başlatalım
	http.ListenAndServe(":8000", nil)
}
