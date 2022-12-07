package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func anasayfa(w http.ResponseWriter, r *http.Request) {
	//html kodumuz
	htmlKodu := "<h1>Merhaba</h1>"

	şablon, _ := template.ParseFiles("sablon.html")

	şablon.Execute(w, htmlKodu)
}
func main() {
	fmt.Println("Sunucu Başladı")

	http.HandleFunc("/", anasayfa)

	http.ListenAndServe(":8000", nil)
}
