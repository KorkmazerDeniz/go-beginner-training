package main

import (
	"log"
	"os"
)

func main() {

	dosya, err := os.Create("html")
	if err != nil {
		log.Fatal(err)
	}

	dosya.Close()
}
