package main

import(
	"fmt"
	"strconv"
)


func main() {
	//Soal 1
var panjangPersegiPanjang string = "8"
var lebarPersegiPanjang string = "5"

var alasSegitiga string = "6"
var tinggiSegitiga string = "7"

panjang, _ := strconv.Atoi(panjangPersegiPanjang)
lebar, _ := strconv.Atoi(lebarPersegiPanjang)
alas, _ := strconv.Atoi(alasSegitiga)
tinggi, _ := strconv.Atoi(tinggiSegitiga)

var luasPersegiPanjang int = panjang * lebar
var kelilingPersegiPanjang int = 2 * panjang * lebar
var luasSegitiga int = int(0.5 * float64(alas) * float64(tinggi))

fmt.Println("Luas Persegi Panjang : ",luasPersegiPanjang)
fmt.Println("Keliling Persegi Panjang : ",kelilingPersegiPanjang)
fmt.Println("Luas Segitiga : ",luasSegitiga)


	//Soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	fmt.Println("")
	fmt.Println("Nilai John :",nilaiJohn)
	if nilaiJohn >= 80 {
		fmt.Println("John mendapatkan nilai A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("John mendapatkan nilai B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("John mendapatkan nilai C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("John mendapatkan nilai D")
	} else {
		fmt.Println("John mendapatkan nilai E")
	}

	fmt.Println("Nilai Doe :",nilaiDoe)
	if nilaiDoe >= 80 {
		fmt.Println("Doe mendapatkan nilai A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("Doe mendapatkan nilai B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("Doe mendapatkan nilai C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("Doe mendapatkan nilai D")
	} else {
		fmt.Println("Doe mendapatkan nilai E")
	}

	//Soal 3
	var tanggal = 3
	var bulan = 3
	var tahun = 2003

	fmt.Println("")
	switch bulan {
	case 1:
		fmt.Println(tanggal, "January", tahun)
	case 2:
		fmt.Println(tanggal, "February", tahun)
	case 3:
		fmt.Println(tanggal, "Maret", tahun)
	case 4:
		fmt.Println(tanggal, "April", tahun)
	case 5:
		fmt.Println(tanggal, "Mei", tahun)
	case 6:
		fmt.Println(tanggal, "Juni", tahun)
	case 7:
		fmt.Println(tanggal, "July", tahun)
	case 8:
		fmt.Println(tanggal, "Agustus", tahun)
	case 9:
		fmt.Println(tanggal, "September", tahun)
	case 10:
		fmt.Println(tanggal, "Oktober", tahun)
	case 11:
		fmt.Println(tanggal, "November", tahun)
	case 12:
		fmt.Println(tanggal, "Desember", tahun)
	}

	//Soal 4
	var kelahiran = 2003
	fmt.Println("")
	fmt.Println("saya kelahiran 2023")

	if kelahiran > 1944 && kelahiran < 1964 {
		fmt.Println("Saya Generasi Boomer")
	}else if kelahiran > 1965 && kelahiran < 1979{
		fmt.Println("Saya Generasi X")
	} else if kelahiran > 1980 && kelahiran < 1994{
		fmt.Println("Saya Generasi X")
	} else if kelahiran > 1995 && kelahiran < 2015{
		fmt.Println("Saya Generasi Z")
	}else{
		fmt.Println("Saya Generasi Baru")
	}
}