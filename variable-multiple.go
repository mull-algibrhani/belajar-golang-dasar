package main

import "fmt"

func main() {
	var (
		nama = "Musliadi"
		usia = 30
		hobi = "Ngoding"
	) 

	fmt.Println(nama, usia, hobi)

	const (
		nama_depan    = "Mull"
		nama_belakang = "algibrhani"
		umur          = 35
	)
}

// variable var yang dideklarasikan tetapi tidak digunakan akan dibaca error
// variable const yang dideklarasikan tetapi tidak digunakan tidak akan dibaca error
// namun variable constant adalah variable yang datanya tidak dapat diubah lagi
