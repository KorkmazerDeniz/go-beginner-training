package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	params := url.Values{
		"q":  {"golang"},
		"hl": {"tr"},
	}

	//URL'imizi aşağıdaki gibi yazalım.
	cevap, hata := http.Get("https://www.google.com.tr/search?" + params.Encode())
	if hata != nil {
		panic(hata)
	}
	defer cevap.Body.Close()

	sayfa, hata := io.ReadAll(cevap.Body)

	if hata != nil {
		panic(hata)
	}

	fmt.Println(string(sayfa))

}
