package main

import (
	"gopkg.in/gomail.v2"
	"log"
)

func main() {

	mesaj := gomail.NewMessage()

	mesaj.SetHeader("From", "ermat.denizkorkmazer@gmail.com")
	mesaj.SetHeader("To", "renault")
	mesaj.SetHeader("Subject", "Önemli!!!")

	mesaj.Attach("udemy.jpg")

	mesaj.SetBody("text/html", "<p>Bu bir paragraf</p><h2 style=\"text-align:center\">Bunu Gördün mü?</h2>")

	dia := gomail.NewDialer("smtp.gmail.com", 587, "ermat.denizkorkmazer@gmail.com", "ermat123")

	err := dia.DialAndSend(mesaj)
	if err != nil {
		log.Fatal(err)
	}

}
