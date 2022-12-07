package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	geciciKlasor, err := os.MkdirTemp("", "udemy")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Gecici Klasor:", geciciKlasor)

	defer os.RemoveAll(geciciKlasor)

	geciciDosya, err := os.CreateTemp(geciciKlasor, "udemykursu")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Gecici Dosya:", geciciDosya.Name())

	defer os.Remove(geciciDosya.Name())

}
