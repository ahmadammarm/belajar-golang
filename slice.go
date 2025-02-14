package main

import (
	"fmt"
	"sort"
	"strconv"
)

// ini akan memicu error karena func main sudah diinisiasi di file lain yaitu file main
func main() {
	var angka []int
	var masukan string

	for {
		fmt.Print("Masukkan angka bebas (atau ketik 'X' untuk keluar dari program): ")
		fmt.Scan(&masukan)

		if masukan == "X" || masukan == "x" {
			break
		}

		ang, err := strconv.Atoi(masukan)
		if err != nil {
			fmt.Println("Masukan tidak valid. Silakan masukkan angka atau ketik 'X' untuk keluar.")
			continue
		}

		angka = append(angka, ang)
		sort.Ints(angka)            

		fmt.Println("Slice yang sudah diurutkan:", angka)
	}

	fmt.Println("Program Selesai")

	// x := [...]int{4, 8, 5}
	// y := x[0:2]
	// z := x[1:3]
	// y[0] = 1
	// z[1] = 3
	//  fmt.Print(x)

}
