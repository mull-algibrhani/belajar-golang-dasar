package main

import "fmt"

func main() {
	// INDEX DI DATA ARRAY
	// INDEX									DATA
	// 0													Mull
	// 1													All
	// 2													Gib
	// 3													Rhani

	var data [4]string
	data[0] = "Mull"
	data[1] = "All"
	data[2] = "Gib"
	data[3] = "Rhani"

	fmt.Println(data[0])
	fmt.Println(data[1])
	fmt.Println(data[2])
	fmt.Println(data[3])
	fmt.Println(data)
	fmt.Println(data[0] + data[1] + data[2] + data[3])
	fmt.Println(data[0], data[1], data[2]+data[3])

	var angka [3]int
	angka[0] = 2
	angka[1] = 5
	angka[2] = 3

	fmt.Println(angka[0], angka[1], angka[2])
	fmt.Println(angka[0] + angka[1] + angka[2])

	var arraykosong [5]int
	fmt.Println(arraykosong)
	arraykosong[3] = 10
	fmt.Println(arraykosong)

	var values = [3]int{20, 10}
	fmt.Println(values)

	// Data array yang tidak di tentukan jumlah datanya di awal
	var values2 = [...]int{
		10, 20, 30, 40, 50,
	}
	fmt.Println(values2)

	// OERASI																								KETERANGAN
	// len(array)																				Untuk mendapatkan panjang array
	// array(index)																		Untuk mendapatkan data di posisi array
	// array[index] = value										Untuk mengubah data di posisi index
	var values3 = [...]int{
		10, 20, 30, 40, 50,
	}

	fmt.Println(values3)

	values3[1] = 8 // mengubah data pada array
	fmt.Println(values3)
	fmt.Println(len(values3))
}
