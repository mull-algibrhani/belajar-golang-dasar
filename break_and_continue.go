package main

import "fmt"

func main() {
	// break digunakan untuk menghentikan perulangan jika kondisi bernilai true
	fmt.Println("BREAK")
	for i := 1; i < 10; i++ {
		if i == 6 {
			break
		}
		fmt.Println("Perulangan ke", i)
	}
	fmt.Println("Program 1 Selesai")
	fmt.Println("\n")
	fmt.Println("Mulai Program 2")
	for i := 1; i < 15; i++ {
		if i == 11 {
			break
		}
		fmt.Println("Perulangan ke", i)
	}
	fmt.Println("Program 2 Selesai")
	fmt.Println("\n")

	// Continue digunakan untuk melanjutkan perulangan jika kondisi bernilai true
	fmt.Println("CONTINUE")
	fmt.Println("Mulai Program 3")
	for i := 0; i < 11; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke", i)
	}
	fmt.Println("Program 3 Selesai")
}
