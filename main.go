package main

// golang memiliki harus memiliki package dan function main sebagai entry point atau eksekusi program

// golang memiliki case sensitive artinya huruf besar dan kecil dapat mempengaruhi program
// package adalah keyword yang digunakan untuk deklarasi package yang digunakan dalam program

// fmt adalah package yang berisi fungsi-fungsi untuk input dan output
// fmt adalah singkatan dari format
import (
	"fmt"
	"time"
	// "go/types"
	// "strconv"
	"math/rand"
	"reflect"
	"strings"

	// Mengimport package lain dari folder yang berbeda menggunakan nama module/package
	"belajar/hello"
	"belajar/aritmatika"
)

// Ingin mengimport package lain, bisa menggunakan import "nama_package"
// Jika ingin mengimport package yang ada di dalam folder, bisa menggunakan import "./nama_folder/nama_package"
// Jika ingin mengimport package yang ada di luar folder, bisa menggunakan import "nama_folder/nama_package"

// import "hello"

// MENJALANKAN PROGRAM GOLANG DI TERMINAL

// Menjalankan menggunakan compiler
// go build belajar.go (untuk mengcompile program)
// Akan menghasilkan file binary belajar yang executable atau bisa dijalankan
// ./belajar (untuk menjalankan program)

// Di dalam go kita bisa mengcompile program dan langsung dijalankan menggunakan go run
// Dalam go run tidak akan menghasilkan file binary yang executable

// Dalam golang, tidak diperbolehkan membuat nama function main lebih dari satu meskipun berada di file yang berbeda
// Jika ada lebih dari satu function main, maka akan terjadi error saat compile program atau build program
// Namun ketika menggunakan go run, maka tidak akan terjadi error dan mencetak yang ada di file yang di run

// Untuk concat string dengan integer bisa menggunakan koma (,) atau menggunakan + dengan fmt.Sprint() yang mengembalikan string

// Penulisan function, jika return value, maka harus dideklarasikan tipe data return valuenya
// Jika tidak ada return value, maka tidak perlu dideklarasikan tipe data return valuenya seperti langsung fmt.Println()

// jika ingin memanggil function lain ke dalam function main, maka function lain tersebut harus me return sesuatu

func duastring(a, b string) (string, string) {
	return a, b
}

func penjumlahan(x int, y int) int {
	return x + y
}

func isPrime(angka int) bool {
	if angka < 2 {
		return false
	}

	for i := 2; i < angka; i++ {
		if angka%i == 0 {
			return false
		}
	}

	return true
}

func findian(ian string) string {

	if strings.HasPrefix(ian, "i") && strings.Contains(ian, "a") && strings.HasSuffix(ian, "n") {
		return "Found!"
	}

	return "Not Found!"
}

func luasLingkaran(jariJari int) int {
	var phi float64 = 3.14
	var luas int = int(phi * float64(jariJari) * float64(jariJari))
	return luas
}

func fizzBuzz(angka int) {
	for i := 1; i <= angka; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func helloWorld() string {
	var hello string = "Hello World"
	return hello
}

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

	// Di golang kita bisa membuat alias atau nama lain dari tipe data
	type NIM string

	var nim NIM = "2131"

	fmt.Println("NIM:", nim)

	var temperature int = 210

	fmt.Println("Temperature:", temperature)

	// Mendefinisikan array
	// di golang, array memiliki panjang yang tetap
	// sintaksnya adalah variable := [jumlah]tipe_data{isi_array}
	// atau bisa menggunakan ... untuk mengisi array tanpa menentukan jumlah array
	var array = [3]int{1, 2, 3}
	var array2 = [...]int{1, 2, 3, 4, 5}

	fmt.Println(array)
	fmt.Println(array2)

	// Memanggil function dari package lain
	// fmt.Println(isPrime(39))

	// Menggunakan pointer di golang dengan tanda &
	var nama_mahasiswa string = "Ahmad Ammar Musyaffa"
	var nama_kampus string = "Universitas Negeri Malang"

	fmt.Println("Nama saya adalah", nama_mahasiswa, "dan saya kuliah di", nama_kampus)

	fmt.Println(&nama_mahasiswa)
	fmt.Println(&nama_kampus)
	fmt.Println(&nama)
	// pointer hanya bisa digunakan di tipe data tertentu seperti int, string, bool, dll
	// pointer tidak bisa digunakan di array, slice, map, dll
	// pointer hanya bisa digunakan di tipe data yang memiliki alamat memori
	// pointer hanya bisa diakses di function yang sama
	// pointer hanya bisa diakses oleh variable, bukan value

	// jika mengunakan * maka akan mengambil value dari pointer
	// jika menggunakan & maka akan mengambil alamat memori dari variable
	// jika menggunakan * dan & maka akan mengambil value dari alamat memori
	fmt.Println(&nim)
	fmt.Println(luasLingkaran(10))

	// constant boleh tidak digunakan, tidak seperti var
	// constant nilainya tidak bisa diubah
	const ammar = "Ammar"
	const (
		adil   = "Adil"
		sholum = "Sholum"
	)

	fmt.Println(ammar, adil, sholum)

	// Konversi tipe data
	// Contoh tipe data di golang
	// int8, int16, int32, int64
	// nilai dari int8 adalah -128 sampai 127
	// nilai dari int16 adalah -32768 sampai 32767
	// nilai dari int32 adalah -2147483648 sampai 2147483647
	// nilai dari int64 adalah -9223372036854775808 sampai 9223372036854775807
	// uint8, uint16, uint32, uint64
	// nilai dari uint8 adalah 0 sampai 255
	// nilai dari uint16 adalah 0 sampai 65535
	// nilai dari uint32 adalah 0 sampai 4294967295
	// nilai dari uint64 adalah 0 sampai 18446744073709551615
	// float32, float64
	// nilai dari float32 adalah 1.18e-38 sampai 3.4e38
	// nilai dari float64 adalah 2.23e-308 sampai 1.8e308
	// string
	// bool
	// jika nilainya melebihi kapasitas tipe data maka akan terjadi overflow atau underflow
	// overflow adalah ketika nilai melebihi kapasitas tipe data dan nilai akan kembali ke nilai terkecil
	// underflow adalah ketika nilai kurang dari kapasitas tipe data dan nilai akan kembali ke nilai terbesar
	var sisi int = 33
	fmt.Println(float32(sisi))

	var nama_byte = "Ammar"
	var byte = nama_byte[2]
	var stringByte = string(byte)
	fmt.Println(stringByte)

	// Latihan soal 1
	var namaAmmar string = "Ammar"
	var umurAmmar int = 21
	var beratBadan float64 = 70.5
	var sudahMenikah string = "Belum menikah"

	fmt.Println("Nama:", namaAmmar)
	fmt.Println("Umur:", umurAmmar)
	fmt.Println("Berat Badan:", beratBadan, "kg")
	fmt.Println("Status:", sudahMenikah)

	// Latihan soal 2
	var mataKuliah string = "Pemrograman Go"
	sks := 3

	fmt.Println("Mata Kuliah:", mataKuliah)
	fmt.Println("Jumlah SKS:", sks)

	// Latihan soal 3
	var angkaSoal float64 = 42.75
	fmt.Println("Nilai floatnya adalah", angkaSoal)
	var angkaInt = int(angkaSoal)
	fmt.Println("Angka integernya adalah", angkaInt)

	// Latihan soal 4
	// var angkaString string = "100"
	// fmt.Println("String sebelum dikonversi adalah", angkaString)
	// angkaInteger, err := strconv.Atoi(angkaString)
	// if err != nil {
	// 	fmt.Println("Error converting string to integer:", err)
	// } else {
	// 	fmt.Println("Sesudah dikonversi menjadi", angkaInteger)
	// }

	fmt.Println(isPrime(21))

	var x int = 10
	var y int

	// tanda * di bawah ini digunakan untuk mengambil value dari pointer
	var ip *int

	ip = &x
	y = *ip
	fmt.Println(y, ip)

	var a int = 100
	var b int = 200

	fmt.Println(&a, *&b)

	// method new digunakan untuk membuat pointer meskipun tidak ada value dalam variable
	angkas := new(int)
	// *angkas = 10
	fmt.Println(angkas)

	type wsas int
	var ws wsas = 10
	fmt.Println(ws)

	// variable scope
	// variable yang dideklarasikan di dalam function hanya bisa diakses di dalam function tersebut
	// variable yang dideklarasikan di luar function bisa diakses di dalam function tersebut
	// variable yang dideklarasikan di luar function bisa diakses di dalam function lain
	// variable yang dideklarasikan di dalam function tidak bisa diakses di dalam function lain

	fmt.Println(helloWorld())

	var angkaSoalPointer int = 10
	var ptr = *&angkaSoalPointer

	fmt.Println("Alamat dari variable tersebut adalah", &angkaSoalPointer)
	fmt.Println("Nilai dari pointer ptr adalah", ptr)

	var nilaiSebelumBerubah int = 50
	var nilaiSetelahBerubah int

	var pointerAngka *int
	pointerAngka = &nilaiSebelumBerubah
	// nilaiSetelahBerubah = *pointerAngka

	fmt.Println(nilaiSebelumBerubah)
	fmt.Println(nilaiSetelahBerubah)
	fmt.Println(pointerAngka)

	// Stack dan Heap memory
	// Stack memory akan menampung variable yang berada di dalam function
	// Sedangkan heap memoery akan menampung variable yang berada di luar function
	// Stack memory akan dihapus ketika function selesai dijalankan

	// cara dialokasi heap memory di golang adalah dengan menggunakan keyword new
	// new akan mengembalikan alamat memori dari variable tersebut
	// new hanya bisa digunakan di tipe data tertentu seperti int, string, bool, dll
	// new tidak bisa digunakan di array, slice, map, dll
	// new hanya bisa digunakan di tipe data yang memiliki alamat memori

	var x1 string = "X"
	xx := "Y"

	fmt.Println(x1, xx)

	// strings method mengimport package strings
	fmt.Println(strings.ToLower("Saya sedang belajar golang"))
	fmt.Println(strings.Count("Aku sedang belajar golang", "a"))

	// Tipe floating
	var trunc float32 = 3.14
	fmt.Println(int(trunc))
	fmt.Println(findian("ian"))

	// x := 7

	// switch {
	// case x > 3:
	// 	fmt.Println("1")
	// case x > 5:
	// 	fmt.Println()
	// }

	fmt.Println("Angka favorit saya adalah", rand.Intn(10)-8)

	hasil := hello.Hello()
	fmt.Println(hasil)

	// Mendeklarasikan tipe data menggunakan package dari reflect dan menggunakan method TypeOf
	nomor := 123
	fmt.Println(reflect.TypeOf(nomor))

	fmt.Println(penjumlahan(12, 21))

	var int int = 12
	var konversi float64 = float64(int)

	fmt.Println(konversi)

	var1, var2, var3 := 1, 2, 3
	var int1, int2, int3 = 1, 2, 3
	fmt.Println(int1, int2, int3, int1+int2+int3)
	fmt.Println(var1, var2, var3, var1+var2+var3)

	w, e := duastring("hello", "world")
	fmt.Println(w, e)

	ab := 5

	switch {
	case ab > 3:
		fmt.Println("World")
	case ab > 4:
		fmt.Println("Hello")
	}

	// Memanggil package time dan salah satu methodnya
	times := time.Now()
	formattedTime := times.Format("Monday, 02 January 2006")
	fmt.Println(formattedTime)
	// fmt.Println(time)
	fizzBuzz(20)

	// Untuk operasi aritmatika harus memiliki tipe data yang sama
	// Misal disini ingin menambahkan kedua bilangan float dan integer
	// Jika ingin hasilnya adalah bilangan float maka harus mengonversi ke float terlebih dahulu
	// Begitu juga jika ingin hasilnya integer, harus dikonversi ke integer terlebih dahulu untuk angka floatnya
	var angkaFloat float64 = 3.14
	var angkaInteger int16 = 100

	var result = angkaFloat + float64(angkaInteger)
	fmt.Println(result)

	var namaString string = "Ammar"
	var namaByte = namaString[3]
	fmt.Println(namaByte)
	var balikString string = string(namaByte)
	fmt.Println(balikString)

	type ammarstring string
	var namaku ammarstring = "Ammar"
	fmt.Println(namaku)

	fmt.Println(aritmatika.Aritmatika(1, 2))
	
	var angkaBaru int8 = 30
	angkaBaru++
	fmt.Println(angkaBaru)

	var angkaBaru2 int8 = 34
	fmt.Println(angkaBaru != angkaBaru2)


	// Latian soal

	// Ada sebuah taman yang memiliki sebuah kolam berbentuk lingkaran
	// Kolam tersebut memiliki jari-jari sekitar 10m
	// Luas dari taman tersebut adalah 20hektar
	// Berapakah luas dari bagian taman selain kolam

	// Ukuran luas taman berdasarkan satuan m^2
	var luas_taman int16 = 2000

	// jari jari kolam dalam satuan meter
	var jari_jari_kolam int16= 10
	var pi float64= 3.14


	var float_luas_taman = float64(luas_taman)
	var float_jari_jari = float64(jari_jari_kolam)

	luas_kolam := pi * float_jari_jari * float_jari_jari

	fmt.Println("Luas kolam adalah", luas_kolam)

	bagian_taman := float_luas_taman - luas_kolam

	fmt.Println("Jadi luas dari bagian taman tanpa kolam tersebut adalah", bagian_taman)

}
