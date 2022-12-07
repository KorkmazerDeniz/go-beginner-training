package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	dosyaOkuma, err := ioutil.ReadFile("ulkeler.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(dosyaOkuma))

	//for i, _ := range dosyaOkuma {
	//	fmt.Println(string(dosyaOkuma[i]))
	//}

	//dosya,err:=os.Open("ulkeler.txt")
	//defer dosya.Close()
	//if err !=nil{
	//	log.Fatal(err)
	//}
	//
	//scanner:=bufio.NewScanner(dosya)
	//
	//for scanner.Scan(){
	//	satir:=scanner.Text()
	//	time.Sleep(2*time.Second)
	//	fmt.Println(satir)
	//
	//}

}
