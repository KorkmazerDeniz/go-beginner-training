package main

import (
	"log"
	"os"
)

//O_RDONLY	SADECE OKUMA
//O_WRONLY	SADECE YAZMA
//O_RDWR	HEM OKUMA HEM YAZMA
//O_APPEND	DOSYA SONUNA EKLEME
//O_CREATE	DOSYA OLUŞTURMA
//O_EXCL	DOSYA YOKSA OLUŞTUR O_CREATE İLE KULLANILIR
//O_SYNC	I/O İşlemleri için Senkron açılır
//O_TRUNC	DOSYA SONUNU KESER

//0000     İzin yok
//0700     Yönetici yazabilir,okuyabilir ve oluşturabilir
//0770     Yönetici ve Grup yazabilir,okuyabilir ve oluşturabilir
//0777     Herkes yazabilir,okuyabilir ve oluşturabilir
//0111     Oluşturma
//0222     Yazma
//0333     Yazma ve Oluşturma
//0444     Okuma
//0555     Okuma ve Oluşturma
//0666     Yazma ve Okuma
//0740     Yönetici yazabilir,okuyabilir ve oluşturabilir; Grup Sadece Okuyabilir; Diğerlerinin İzni Yok

//Birden fazla şekilde de kullanılabilir O_WRONLY|O_RDONLY

func main() {

	dosya, err := os.OpenFile("veri.txt", os.O_WRONLY, 0000)
	defer dosya.Close()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Dosya Açıldı")

	//dosya.WriteString("Abdulhamid Kaan")

	//yaziSlice:=[]byte("Merhabalar Udemy ")
	//
	//dosyaYazma,err:=dosya.Write(yaziSlice)
	//if err !=nil{
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("Dosyaya %d Byte Boyutunda Yazı Yazıldı",dosyaYazma)

	sehirler := []string{"Malatya", "Ankara", "Çanakkale", "İzmir", "Kütahya"}

	for _, sehir := range sehirler {
		_, err := dosya.WriteString(sehir + "\n")
		if err != nil {
			log.Fatal(err)
		}

	}

}
