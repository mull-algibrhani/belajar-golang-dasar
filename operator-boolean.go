package main

import "fmt"

func main() {

	// OPERATOR BOOLEAN  KETERANGAN
	// &&   AND										DAN
	// ||   OR											ATAU
	// !    NOT										TIDAK/KEBALIKAN

	//OPERASI &&
	// NILAI 1     OPERATOR     NILAI 2     HASIL
	// TRUE								&&											TRUE								TRUE
	// TRUE								&&											FALSE							FALSE
	// FALSE							&&											TRUE							 FALSE
	// FALSE							&&											FALSE							FALSE

	//OPERASI ||
	// NILAI 1     OPERATOR     NILAI 2     HASIL
	// TRUE								||											TRUE								TRUE
	// TRUE								||											FALSE							TRUE
	// FALSE							||											TRUE							 TRUE
	// FALSE							||											FALSE							FALSE

	//OPERASI !
	// OPERATOR     NILAI 2     HASIL
	// !											 TRUE								FALSE
	// !											 FALSE							TRUE

	var nilaiUjian = 90
	var nilaiAbsen = 70
	var lulusUjian = nilaiUjian > 80
	var luluAbsen = nilaiAbsen > 80

	var lulus1 = lulusUjian && luluAbsen
	var lulus2 = lulusUjian || luluAbsen
	fmt.Println(lulus1)

	if lulusUjian && luluAbsen {
		fmt.Println("OPERATOR && AND")
		fmt.Println("LULUS")
	} else {
		fmt.Println("OPERATOR && AND")
		fmt.Println("TIDAK LULUS")
	}

	fmt.Println("\n")
	fmt.Println(lulus2)
	if lulusUjian || luluAbsen {
		fmt.Println("OPERATOR || OR")
		fmt.Println("LULUS")
	} else {
		fmt.Println("OPERATOR || OR")
		fmt.Println("TIDAK LULUS")
	}

	fmt.Println("\n")
	if !luluAbsen {
		fmt.Println("OPERATOR ! NOT")
		fmt.Println("TIDAK LULUS")
	} else {
		fmt.Println("OPERATOR ! NOT")
		fmt.Println("LULUS")
	}

}
