package main

import "fmt"

func declareData() {
	var i = 21

	fmt.Printf("nilai i = %d\n", i)
}

func declareTipe() {
	var i = 22
	fmt.Printf("Tipe data variable i = %T\n", i)
}
func declareTanda() {
	fmt.Printf("Tanda = %%\n")
}

func declareBool() {
	var condition bool = true
	fmt.Printf("nilai boolean j = %t\n", condition)
}
func declareUnic() {
	r1 := '\u042f'
	fmt.Printf("Unicode Russia = %c\n", r1)
}
func base10() {
	var i = 21
	fmt.Printf("Nilai base 10 = %d\n", i)
}
func base8() {
	var i = 21
	fmt.Printf("nilai base 8 = %o\n", i)
}
func declaref() {
	fmt.Printf("Nilai base 16 : %x\n", 15)
}
func declareF() {
	fmt.Printf("Nilai base 16 : %X\n", 15)
}
func unikara() {
	r1 := '\u042f'
	fmt.Printf("Karakter Я dalam Unicode: %U\n", r1)
}
func declare64() {
	var k float64 = 123.456
	fmt.Printf("float = %f\n", k)
	fmt.Printf("float cientific :%E", k)
}
func main() {
	declareData()
	declareTipe()
	declareTanda()
	declareBool()
	declareUnic()
	base10()
	base8()
	declaref()
	declareF()
	unikara()
	declare64()
}
