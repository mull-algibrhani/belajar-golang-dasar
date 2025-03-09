package main

import "fmt"

func main() {
	var name string
	name = "Rinra"

	// IF Statement
	if name == "mull" {
		fmt.Println("Sama")
	} else {
		fmt.Println("IF Statement")
		fmt.Println("Tidak Sama")
		fmt.Println("\n")
	}

	// ELSE IF Statement
	fmt.Println("ELSE IF Statement")
	if name == "Mull" {
		fmt.Println("Nama kamu Mull")
	} else if name == "mull" {
		fmt.Println("Nama kamu mull")
	} else if name == "Andre" {
		fmt.Println("Nama kamu Andre")
	} else if name == "Rinra" {
		fmt.Println("Nama kamu Rinra")
	} else {
		fmt.Println("Hai, Kenalan Dong")
		fmt.Println("\n")
	}

	//  IF sort statment
	if length := len(name); length > 4 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Printf("Nama sudah benar")
	}

}
