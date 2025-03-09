package main

import "fmt"

func main() {
	var month = [...]string{"januari", "februari", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november", "desember"}
	var slice1 = month[4:6]
	var slice2 = month[4:]
	var slice3 = month[:6]
	var slice4 = month[:]

	fmt.Println(slice1)
	fmt.Println(slice1[0])
	fmt.Println(slice1[1])

	fmt.Print("\n")
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice4[7])
	fmt.Print("\n")

	// Mengubah data pada slice
	var days = [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daySlice1 := days[5:] // sabtu, minggu

	fmt.Println(days)
	fmt.Println(daySlice1)
	fmt.Print("\n")

	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(daySlice1)
	fmt.Print("\n")
	fmt.Println(days)

	//menggunakan append untuk menambah data array
	daySlice2 := append(daySlice1, "senin lagi", "minggu lagi")
	fmt.Print("\n")
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)
	fmt.Print("\n")

	// menggunakan make slice
	var slicebaru []string = make([]string, 2, 5)
	slicebaru[0] = "Mull"
	slicebaru[1] = "Algibrhani"
	// slicebaru[2] = "Magassing" // erro karena panjang array sudah ditentukan sebanyak 2
	fmt.Println(slicebaru)
	fmt.Println(len(slicebaru))
	fmt.Println(cap(slicebaru))
	fmt.Print("\n")

	slicebaru2 := append(slicebaru, "Musliadi")
	fmt.Println(slicebaru2)
	fmt.Println(len(slicebaru2))
	fmt.Println(cap(slicebaru2))
	fmt.Print("\n")

	slicebaru2[0] = "Mus"
	fmt.Println(slicebaru)
	fmt.Println(slicebaru2)

	// Menggunakan copy slice
	darislice := days[:]
	keslice := make([]string, len(darislice), cap(darislice))
	fmt.Println(keslice)
	copy(keslice, darislice)

	fmt.Print("\n")
	fmt.Println(darislice)
	fmt.Println(keslice)

	// perbedaaan pembuatan type data array dan slice

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Print("\n")
	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
