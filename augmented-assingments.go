package main

import "fmt"

func main() {
	// Operasi Matematika					Augmented Assigments
	// a = a + 10													a += 10
	// a = a - 10													a -= 10
	// a = a * 10													a *= 10
	// a = a / 10													a /= 10
	// a = a % 10													a %= 10
	var a = 10
	var b = 5
	var c = 4
	var d = 6
	var e = 7
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	a += 10 // a = 10 + 10 = 20
	b -= 2  // b = 5 - 2 = 3
	c *= 2  // c = 4 * 2 = 8
	d /= 3  // d = 6 / 3 = 2
	e %= 5  // e = 7 % 5 = 2
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
