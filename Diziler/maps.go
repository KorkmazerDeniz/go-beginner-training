package main

import "fmt"

// MAPS
func main() {

	//Standart oluşturma yöntemi

	var ogrenci map[int]string = map[int]string{
		10: "Deniz",
		21: "Ferhat",
		39: "Sercan",
		89: "Merve",
	}

	fmt.Println(ogrenci)

	fmt.Println(ogrenci[39])

	ogrenci[21] = "Ferhat"

	fmt.Println(ogrenci)

	//key:int  value:string
	plaka := make(map[int]string)

	plaka[44] = "İstanbul"
	plaka[35] = "İzmir"
	plaka[01] = "Sivas"

	fmt.Println(plaka)

	fmt.Println(plaka[35])

	//key:string value:string

	unvanlar := make(map[string]string)

	unvanlar["DR"] = "DOKTOR"
	unvanlar["PRF"] = "PROFESOR"
	unvanlar["DOC"] = "DOCENT"

	fmt.Println(unvanlar["PRF"])

	a := unvanlar["DR"]

	fmt.Println(a)

	//silme

	delete(unvanlar, "DR")

	fmt.Println(unvanlar)

	fmt.Println(len(unvanlar))

	for key, value := range unvanlar {
		fmt.Printf("Kısaltılması:%v Tam Hali:%v\n", key, value)
	}

}
