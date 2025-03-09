package main

import "fmt"

func main() {
	// Statement switch digunakan untuk melakukan pengecekan sederhana dengan hanya 1 buah variable
	// Didalam switch tidak dapat menggunakan operator boolean seperti && || dan !

	// Switch statement
	fmt.Println("Switch  statement")

	grade := "C"

	switch grade {
	case "A":
		fmt.Println("Kualitas sangat baik")
	case "B":
		fmt.Println("Kualitas baik")
	case "C":
		fmt.Println("Kualitas cukup baik")
	case "D":
		fmt.Println("Kualitas kurang baik")
	case "E":
		fmt.Println("Kualitas tidak baik")
	default:
		fmt.Println("Kualitas tidak valid")

	}
	fmt.Println("\n")

	// Switch short statement
	fmt.Println("Switch short statement")
	switch length := len(grade); length > 0 {
	case true:
		fmt.Println("Nilai Grade Ada")
	case false:
		fmt.Println("Nilai Grade Tidak Ada")
	}
	fmt.Println("\n")

	// switch statement tanpa kondisi
	fmt.Println("switch statement tanpa kondisi")

	name := "Mull"
	length := len(name)
	switch {
	case length == 8:
		fmt.Println("Panjang Karakter adalah 8")
	case length == 4:
		fmt.Println("Panjang Karakter adalah 4")
	case length == 5:
		fmt.Println("Panjang Karakter adalah 5")
	default:
		fmt.Println("Panjang Karakter diluar jangkauan")
	}

}
