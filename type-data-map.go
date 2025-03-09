package main

import "fmt"

func main() {
	var biodata map[string]string = map[string]string{}
	biodata["nama"] = "Mull"
	biodata["alamat"] = "Moti"

	nilaiUjian := map[string]float64{
		"Matematika": 2.5,
		"Bahasa":     2.5,
	}

	fmt.Println("Nama : ", biodata["nama"])
	fmt.Println("Nilai Ujian Matematika =", nilaiUjian["Matematika"])
	fmt.Println("Nilai Ujian Bahasa =", nilaiUjian["Bahasa"])
	fmt.Println("\n")

	// Function Map
	// Operasi																						// Keterangan
	// len(map)																					// Untuk Mendapatkan jumlah data di map
	// map[key]																					// mengambil data di map dengan key
	// map[key] = value													// mengubah data di map dengan key
	// make(map[TypeKey]TypeValue)  // membuat map baru
	// delete(map, key)													// menghapus data di map dengan key
	// len(map[key])																// untuk mendapatkan jumlah panjang data di map

	// Make
	book := make(map[string]string)
	book["title"] = "Buku Filsafat"
	book["author"] = "Aris thoteles"
	book["wrong"] = "Ups"
	fmt.Println(book)

	// len
	fmt.Println(len(book))

	// delete
	delete(book, "wrong")
	fmt.Println(book)

	// len
	fmt.Println(len(book))

	// map key
	fmt.Println(book["title"])

	// map[key] = value
	book["title"] = "Buku Sejearah"
	fmt.Println(book)

	// len(map[key])
	fmt.Println(len(book["author"]))
}
