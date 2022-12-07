package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func checkErr4(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	db, err := sql.Open("sqlite3", "filmler.db")
	defer db.Close()
	checkErr4(err)

	//GUNCELEME

	guncelle, err := db.Prepare("UPDATE oyuncular SET yas=? WHERE ad=?") //stmt
	checkErr4(err)
	res, err := guncelle.Exec(99, "Orhan")

	affect, err := res.RowsAffected()
	//affect, err := res.LastInsertId()
	checkErr4(err)
	fmt.Println("İşlem Yapılan Satır Sayısı", affect)

}
