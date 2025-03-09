package main

import "fmt"

func main() {

	                            
	// variable constant walaupun sudah di deklarasikan namun tidak dipakai, tidak akan error. berbeda dengan var
	const nama_depan = "Mull"
	const usia = 20
	const hobi string = "membaca"

	const nilai = 2.5

	fmt.Println(nama_depan)
	fmt.Println(usia)
	fmt.Println(hobi)

	//error
	// nama_depan = "Musliadi"
	// usia = 32
	// hobi = "memancing"

	//error
	// const alamat string

}
