package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" //init
	"log"
)

func checkErr(err error) { //OVERLOADING
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	//_, err := os.Create("data1.db")  KULLANMANIZA GEREK YOK
	//checkErr(err)
	db, err := sql.Open("sqlite3", "filmler.db")
	defer db.Close()
	checkErr(err)

	_, err = db.Exec("CREATE TABLE notlar(id INT PRIMARY KEY,notdurumu TEXT,isim TEXT)")
	checkErr(err)

}
