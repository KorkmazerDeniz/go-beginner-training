package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func checkErr3(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	db, err := sql.Open("sqlite3", "filmler.db")
	defer db.Close()
	checkErr3(err)

	oyuncuAd := []string{"Şener Şen", "Kemal Sunal"}
	oyuncuId := []int{4, 5}
	oyuncuYas := []int{74, 68}

	for i := 0; i < len(oyuncuId); i++ {
		veriekle, err := db.Prepare("INSERT INTO oyuncular(id,ad,yas)VALUES(?,?,?)")
		checkErr3(err)
		veriekle.Exec(oyuncuId[i], oyuncuAd[i], oyuncuYas[i])
	}

}
