package main

import "fmt"

func main() {
	// konversi type data int
	var nilai32 int32 = 32769
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	//konversi tipe data string
	var name = "Mull Algibrhani"
	//mengambil index 0 dengan return byte
	var e = name[0]
	//konversi index 0 dari byte ke text/string
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
	//mengambil index ke 5 dan langsung dikonversi ke string
	fmt.Println(string(name[5]))
}
