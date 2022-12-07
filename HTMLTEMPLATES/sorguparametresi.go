func main() {
params := url.Values{
"q":  {"golang"}, //aranacak metin
"hl": {"tr"}, //örnek olarak Türkçe sonuç vermesi için
}
fmt.Println("https://www.google.com.tr/search?" + params.Encode())
}
