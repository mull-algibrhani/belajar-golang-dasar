package main

import "fmt"

func main() {
	// jenis operator matematika
	// + = Penjumlahan
	// - = Pengurangan
	// * = Perkalian
	// / = Pembagian
	// % = Sisa Pembagian

	var a int
	var b int
	var c int
	var d int
	a = 10
	b = 5
	c = 2
	d = 8

	fmt.Println(a, "+", b, "=", a+b)
	fmt.Println(d, "-", c, "=", d-c)
	fmt.Println(c, "x", a, "=", c*a)
	fmt.Println(b, "/", c, "=", b/c)
	fmt.Println(b, "%", c, "=", b%c)

}
