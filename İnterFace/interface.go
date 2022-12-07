package main

import "fmt"

/*
type interface_adi interface{

	// Method imzaları
	func1() int
	func2() float64

}

*/

/*



type A struct{
}
func(x A)Read(){
}

type B struct{
}
func(y B)Read(){
}

func callRead(a A){
     a.Read()
}

func callRead2(b B){
     b.Read()
}



*/

type yuruyen interface {
	Yuru() string
}

type Insan string
type Kedi string

func (i Insan) Yuru() string {
	return "Ben bir insanım ve yürüyorum"
}

func (k Kedi) Yuru() string {
	return "Ben bir kediyim ve yürüyorum"
}

// Interface'den oluşturulan Func
func yurume(y yuruyen) {
	fmt.Println(y.Yuru())
}

func main() {

	var insan Insan
	var kedi Kedi

	yurume(insan)
	yurume(kedi)

}
