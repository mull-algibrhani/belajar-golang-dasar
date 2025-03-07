package main

import "fmt"

func main() {
	// OPERATOR     KETRERANGAN
	// >												Lebih dari
	// <												Kurang dari
	// >=											Lebih dari sama dengan
	// <=											Kurang dari sama dengan
	// ==											Sama dengan
	// !=											Tidak sama dengan
	var a = 5
	var b = 3
	var name1 = "Mull"
	var name2 = "Mull"

	var result1 = a < b          // false
	var result2 = name1 == name2 // true
	var result3 = a != b         // true
	var result4 = a > b          // true
	var result5 = a == b         // false
	var result6 = b <= a         // true
	var result7 = b >= a         // false

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)
	fmt.Println(result7)

}
