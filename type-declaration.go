package main

import "fmt"

func main() {

	type NoKTP string

	var ktpMull NoKTP = "7303020405"
	var contoh string = "3334445555"
	var contoKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpMull)
	fmt.Println(contoKtp)

}
