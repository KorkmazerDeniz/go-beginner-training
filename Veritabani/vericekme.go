package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func checkErr6(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	db, err := sql.Open("sqlite3", "filmler.db")
	defer db.Close()
	checkErr6(err)

	satirlar, err := db.Query("SELECT * FROM filmler")
	checkErr6(err)

	var id int
	var filmAd string
	var filmKategori string
	var yonetmenAd string
	var yapimyili int
	var butce int

	for satirlar.Next() {
		err = satirlar.Scan(&id, &filmAd, &filmKategori, &yonetmenAd, &yapimyili, &butce)
		checkErr6(err)

		//fmt.Println(id)
		//fmt.Println(filmAd)
		//fmt.Println(butce)

		fmt.Println("ID:", id, "Film Adı:", filmAd, "Film Kategori:", filmKategori, "Yönetmen:", yonetmenAd, "Yapım Yılı:", yapimyili, "Bütçe", butce/2)

	}

}
