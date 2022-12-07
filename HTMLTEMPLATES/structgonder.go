package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type bilgi struct {
	Başlık string
	İçerik template.HTML
}

func anasayfa(w http.ResponseWriter, r *http.Request) {
	sayfaBilgi := bilgi{
		Başlık: "Sayfa Başlığı",
		İçerik: "<b>sayfa içeriği</b>",
	}
	şablon, _ := template.ParseFiles("sablon.html")
	şablon.Execute(w, sayfaBilgi)
}
func main() {
	fmt.Println("Sunucu Başladı")
	http.HandleFunc("/", anasayfa)
	http.ListenAndServe(":8000", nil)
}
