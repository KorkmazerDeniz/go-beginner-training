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
	checkErr4(err)
	db.Exec("DROP TABLE notlar")

}
