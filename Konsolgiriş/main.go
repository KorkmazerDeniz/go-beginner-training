package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	//tarayici := bufio.NewScanner(os.Stdin)

	//tarayici.Scan()

	//verigirisi := tarayici.Text()

	//fmt.Printf("Bunu yazdım:%s,", verigirisi)

	scanner:=bufio.NewScanner(os.Stdin)

	fmt.Print("Hangi yıl doğdunuz:")

	scanner.scan()

	verigirisi,_:=strconv.ParseInt(scanner.Text(),base:10,bitSize:64)

	fmt.pritf("Şuanda %d yaşındasın.",2021-verigirisi)

}
