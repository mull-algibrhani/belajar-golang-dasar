package main

import "fmt"

func main() {

	var nama string
	var usia int

	nama = "Musliadi"
	usia = 30

	fmt.Println("Nama = ", nama)
	fmt.Println("Usia = ", usia)

	nama = "Mull"
	usia = 35
	//  usia = "35"
	//error jika variable usia yang awalnya adalah int diubah menjadi string

	fmt.Println("Nama = ", nama)
	fmt.Println("Usia = ", usia)

	//contoh lain deklarasi variable dengan := tanpa awalan var
	hobi := "Membaca"
	jumlah_anak := 1
	fmt.Println("Hobi = ", hobi)
	fmt.Println("Jumlah anak = ", jumlah_anak)

	// deklarasi variable dengan := tidak berlaku jika variable lebih awal ditentukan tipe datanya seperti contoh dibawah ini
	// alamat := string
	// alamat = "Baji Minasa"
	// fmt.Println("Alamat = ", alamat)
}

// variable di golang hanya bisa menampung jenis tipe data yang sama
// isi data pada variable hanya dapat diubah dengan tipe data yang sama
// variable yang dideklarasikan tetapi tidak digunakan akan dibaca error
