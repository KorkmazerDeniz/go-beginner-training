package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func checkErr2(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {

	db, err := sql.Open("sqlite3", "filmler.db")
	defer db.Close()
	checkErr2(err)
	//db.Exec("DROP TABLE notlar")
	db.Exec("CREATE TABLE IF NOT EXISTS oyuncular(id INT PRIMARY KEY,ad TEXT,yas INT )")
	oyuncuAd := "Brad Pitt"
	oyuncuYas := 60
	oyuncuId := 2

	veriekle, err := db.Prepare("INSERT INTO oyuncular(id,ad,yas) VALUES(?,?,?)")
	checkErr2(err)
	veriekle.Exec(oyuncuId, oyuncuAd, oyuncuYas)
}
