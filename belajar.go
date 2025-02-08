// golang memiliki harus memiliki function main sebagai entry point atau eksekusi program
package main

// golang memiliki case sensitive artinya huruf besar dan kecil dapat mempengaruhi program
// package adalah keyword yang digunakan untuk deklarasi package yang digunakan dalam program

// fmt adalah package yang berisi fungsi-fungsi untuk input dan output
// fmt adalah singkatan dari format
import "fmt"



// MENJALANKAN PROGRAM GOLANG DI TERMINAL

// Menjalankan menggunakan compiler
// go build belajar.go (untuk mengcompile program)
// Akan menghasilkan file binary belajar yang executable atau bisa dijalankan
// ./belajar (untuk menjalankan program)


// Menjalankan menggunakan interpreter
// Menggunakan interpreter tidak perlu mengcompile program
// go run belajar.go (untuk menjalankan program)


// Dalam golang, tidak diperbolehkan membuat nama function main lebih dari satu meskipun berada di file yang berbeda
// Jika ada lebih dari satu function main, maka akan terjadi error saat compile program atau build program
// Namun ketika menggunakan go run, maka tidak akan terjadi error dan mencetak yang ada di file yang di run

// Untuk concat string dengan integer bisa menggunakan koma (,) atau menggunakan + dengan fmt.Sprint() yang mengembalikan string



// function main akan dieksekusi saat program dijalankan
func main() {
	// Println adalah fungsi untuk menampilkan output ke layar
	// var name string = "Ahmad Ammar"

	// var umur int = 21

	// var isMenikah bool = false

	// fmt.Println(name, "berumur", umur)
	// fmt.Println(12, "asdwq")
	// fmt.Println("Menikah:", isMenikah)

	// Latihan soal

	// var angka1 int = 12
	// var angka2 int = 20

	// var hasil int = angka1 + angka2

	// fmt.Println("hasil dari penjumlahan angka", angka1, "dan", angka2, "adalah", hasil)

	// var name string = "Ammar"
	
	// Mengambil index dalam string namun dalam bentuk byte, bukan tipe data string
	// fmt.Println(name[2])
	
	// Variable var bisa diubah nilainya (seperti var dalam js)
	// var nim string = "2131"
	// nim = "313123"

	// fmt.Println(nim)

	// Variable bisa dideklarasikan tanpa tipe data
	// var nama = "Ahmad Ammar"
	// nama = "Ammar"

	// fmt.Println(nama)

	// Membuat variable bisa menggunakan sintaks :=
	// Jika membuat variable dengan cara := maka tidak boleh dideklarasikan lagi di variable berikutnya
	// Karena penggunaan := hanya untuk pendeklarasian pertama saja

	// Semua variable harus digunakan dan tidak boleh tidak digunakan
	// nama_lengkap := "Ahmad Ammar Musyaffa"
	// nama_lengkap = "Ammar"

	// fmt.Println(nama_lengkap)

	var nama string = "Ammar"
	var umur int = 22
	var tinggi int = 170

	fmt.Println("Nama :", nama)
	fmt.Println("Umur :", umur, "tahun")
	fmt.Println("Tinggi :", tinggi, "cm")

	var namaDepan string = "Ahmad"
	namaBelakang := "Ammar"

	var namaLengkap string = namaDepan + " " + namaBelakang

	fmt.Println("Nama Lengkap:", namaLengkap)

	var jumlah int = 10

	fmt.Println("Nilai awal:", jumlah)

	jumlah = jumlah + 5

	fmt.Println("Nilai setelah ditambah 5:", jumlah)
}