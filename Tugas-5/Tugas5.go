package main

import (
	"fmt"
	"strings"
)

//Soal 1
func LuasPersiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

func KelilingPersegiPanjang(panjang int, lebar int) int {
	return 2 * (panjang * lebar)
}

func VolumeBalok(panjang int, lebar int, tinggi int) int {
	return panjang * lebar * tinggi
}

//Soal 2
func introduce(nama, jenis, hobby, umur string) string {
	return "Pak " + nama + " adalah seorang " + hobby + " yang berusia " + umur + " tahun"
}

//soal 3
func buahFavorit(nama string, buah ...string) string {
	buahFavoritStr := strings.Join(buah, "\", \"")
	return fmt.Sprintf("Halo, nama saya %s dan buah favorit saya adalah \"%s\".", nama, buahFavoritStr)
}

//Soal 4
var dataFilm = []map[string]string{}
func tambahDataFilm() func(string, string, string, string) {
	return func(judul, durasi, genre, tahun string) {
		film := map[string]string{
			"judul":  judul,
			"durasi": durasi,
			"genre":  genre,
			"tahun":  tahun,
		}
		dataFilm = append(dataFilm, film)
	}
}



func main() {
	//Soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := LuasPersiPanjang(panjang, lebar)
	keliling := KelilingPersegiPanjang(panjang, lebar)
	volume := VolumeBalok(panjang, lebar, tinggi)

	fmt.Println("Luas Persegi Panjang = ", luas)
	fmt.Println("Keliling Persegi Panjang = ", keliling)
	fmt.Println("Volume Balok = ", volume)
	fmt.Println("")

	//Soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah)
	fmt.Println(john)
	fmt.Println("")

	//Soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	fmt.Println("")

	//soal 4
	tambahDataFilm := tambahDataFilm()

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}
