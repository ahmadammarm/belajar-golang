package variable

import "fmt"

func variable() {
	// Variable adalah tempat untuk menyimpan data
	// Variable di Go bisa dideklarasikan dengan tipe data atau tanpa tipe data
	// Variable di Go bersifat static, artinya tipe data variable tidak bisa diubah
	// Variable di Go bisa dideklarasikan dengan menggunakan keyword var atau :=

	// Deklarasi variable dengan var dan tipe data
	var nama string = "Ahmad Ammar"
	fmt.Println(nama)

	// Deklarasi variable tanpa tipe data
	var nim = "2131"
	fmt.Println(nim)

	// Deklarasi variable dengan :=
	prodi := "Teknik Informatika"
	fmt.Println(prodi)

	// Perbedaan antara var dan := adalah var bisa dideklarasikan ulang di variable berikutnya
	// Sedangkan := tidak bisa dideklarasikan ulang di variable berikutnya

	// Contoh
	var nama1 string = "Ammar"
	nama1 = "Ahmad Ammar"
	fmt.Println(nama1)

	nama2 := "Mumus"
	// nama2 := "Mumus" // Error
	// Jika tidak ingin error, maka gunakan cara berikut
	nama2 = "Mumus"
	fmt.Println(nama2)

	// := adalah cara yang lebih singkat dan mudah untuk membuat variable
}