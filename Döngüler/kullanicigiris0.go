package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	k_adi := "Deniz"
	k_sifre := "12345"

	giris_hakki := 3

	fmt.Println("****LOGIN****")

	for true {

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Kullanıcı Adını Giriniz:")

		scanner.Scan()

		giris_kadi := scanner.Text()

		fmt.Print("Kullanıcı Şifresini Giriniz:")

		scanner.Scan()

		giris_ksifre := scanner.Text()

		if k_adi != giris_kadi && k_sifre != giris_ksifre {
			fmt.Println("Kullanıcı Adınız ve Şifreniz Yanlış")
			giris_hakki--
		} else if k_adi == giris_kadi && k_sifre != giris_ksifre {
			fmt.Println("Şifreniz Yanlıştır.")
			giris_hakki--
		} else if k_adi != giris_kadi && k_sifre == giris_ksifre {
			fmt.Println("Kullanıcı Adınız Yanlıştır.")
			giris_hakki--
		} else {
			fmt.Println("Hoşgeldiniz")
			break
		}

		if giris_hakki == 0 {
			fmt.Println("Giris Hakkınız Sona Erdi")
			break
		}

	}

}
