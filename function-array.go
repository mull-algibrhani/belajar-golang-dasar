package main

import "fmt"

func main() {
	// OERASI																								KETERANGAN
	// len(array)																				Untuk mendapatkan panjang array
	// array(index)																		Untuk mendapatkan data di posisi array
	// array[index] = value										Untuk mengubah data di posisi index
	var values = [...]int{
		10, 20, 30, 40, 50,
	}

	fmt.Println(values)

	values[1] = 8
	fmt.Println(values)
}
