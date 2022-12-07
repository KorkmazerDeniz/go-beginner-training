package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func anasayfa(w http.ResponseWriter, r *http.Request) {
	şablon, _ := template.ParseFiles("sablon.html")

	//şablona değer göndermeyeceğimiz için burası nil olsun.
	şablon.Execute(w, nil)
}
func main() {
	fmt.Println("Sunucu Başladı")
	http.HandleFunc("/", anasayfa)
	http.ListenAndServe(":8000", nil)
}
