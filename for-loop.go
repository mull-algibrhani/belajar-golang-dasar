package main

import "fmt"

func main() {

	counter := 1

	fmt.Println("Perulangan For dengan init ")
	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	fmt.Println("\n")
	fmt.Println("Perulangan For dengan post statement")
	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke ", i)
	}

	fmt.Println("\n")
	fmt.Println("Perulangan For pada data array dengan manual")
	siswa := [...]string{"Awal", "Indah", "Akbar"}
	no := 1
	for i := 0; i < len(siswa); i++ {
		fmt.Println(no, siswa[i])
		no++
	}

	fmt.Println("\n")
	fmt.Println("Perulangan For dengan range pada data array tanpa manual")
	for index, value := range siswa {
		fmt.Println(index, value)
	}

	fmt.Println("\n")
	fmt.Println("Perulangan For dengan range pada data array tanpa manual tanpa index")
	for _, value := range siswa {
		fmt.Println(value)
	}

	fmt.Println("\n")
	fmt.Println("Perulangan For dengan range pada data map")
	stokbarang := make(map[string]int)
	stokbarang["Pulpen"] = 10
	stokbarang["Buku"] = 5
	stokbarang["Tas"] = 2
	stokbarang["Spidol"] = 7
	fmt.Println("jumlah barang = ", len(stokbarang))
	count := 1
	for key, value := range stokbarang {
		fmt.Println("No.", count, "||", key, "||", "Stok ||", value)
		count++
	}
}
