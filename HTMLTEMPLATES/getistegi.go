package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//Google anasayfasına GET isteğinde bulunalım
	//ve dönen cevabı cevap değişkenine atayalım
	cevap, hata := http.Get("https://www.google.com.tr")
	if hata != nil {
		//Ya sev ya terket dedik ve paniği bastık.
		panic(hata)
	}
	//Body'i kapatmayı unutmayalım
	defer cevap.Body.Close()
	//Bu esnada Body'den cevabı çekip sayfa değişkenine alıyoruz.
	//Çünkü bunu yapmadan okunabilir bir sonuç alamayız.
	sayfa, hata := io.ReadAll(cevap.Body)
	if hata != nil {
		panic(hata)
	}
	//sayfa değişkeni byte dizisi tipinde olduğu için
	//string tipine cevirip okuyalım
	fmt.Println(string(sayfa))
	//Çıktımız Google TR anasayfasının kaynak kodları olacaktır.
}
