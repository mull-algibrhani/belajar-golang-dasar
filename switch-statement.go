package main

import "fmt"

func main() {
	// Statement switch digunakan untuk melakukan pengecekan sederhana dengan hanya 1 buah variable
	// Didalam switch tidak dapat menggunakan operator boolean seperti && || dan !
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

	// Switch short statement

	switch length := len(grade); length > 0 {
	case true:
		fmt.Println("Nilai Grade Ada")
	case false:
		fmt.Println("Nilai Grade Tidak Ada")
	}

}
