package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {

	db, err := sql.Open("sqlite3", "filmler.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	guncelle, err := db.Prepare("UPDATE filmler SET butce=?")
	if err != nil {
		log.Fatal(err)
	}

	tx, BGerr := db.Begin()
	if BGerr != nil {
		log.Fatal(BGerr)
	}

	_, txerr := tx.Stmt(guncelle).Exec(500)
	if txerr != nil {
		fmt.Println("Geri alınıyor...ROLLBACK")
		tx.Rollback()
	} else {
		fmt.Println("Commit Edildi")
		tx.Commit()
	}

}
