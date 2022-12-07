package main

import "fmt"

func main(){

harf:="z"

switch harf{
case "a","e","i","o","u":
fmt.Println("Sesli Harf")
default:
fmt.Println("Sessiz Harf")
}




//Koşullu Switch



sayi:=5

switch{
case sayi==5:
fmt.Println("Sayı 5'e eşittir.")
fallthrough
case sayi<10:
fmt.Println("Sayı 10'dan küçüktür.")
default:
fmt.Println("Tanımsız")
}


*/

switch sayi:=25;{
case sayi>=0 && sayi<=10:
fmt.Println("Sayı 0'dan büyük 10'dan küçük yada eşittir.")
case sayi>=11 && sayi<=30:
fmt.Println("Sayı 11'den büyük ve eşit yada 30'dan küçük yada eşit")
default:
fmt.Println("Aralık tanımsız")
}

}