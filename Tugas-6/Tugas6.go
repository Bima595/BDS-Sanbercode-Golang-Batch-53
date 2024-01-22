package main

import (
	"fmt"
)

//Soal 2
func introduce(sentence *string, name, gender, occupation string, age string) {
	*sentence = "Pak " + name + " adalah seorang " + occupation + " yang berusia " + age + " tahun"
	if gender == "perempuan" {
		*sentence = "Bu " + name + " adalah seorang " + occupation + " yang berusia " + age + " tahun"
	}
}

//Soal 3
func buah(){
	var buah = []string{"Jeruk", "Semangka", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis","Alpukat"}
	var listbuah *string
	for i, namaBuah := range buah {
		listbuah = &namaBuah
		fmt.Printf("%d. %s\n", i+1, *listbuah)
	}
	fmt.Println("")
}


//Soal 4
var dataFilm = []map[string]string{}

func tambahDataFilm(listFilm *[]map[string]string) func(string, string, string, string) {
	return func(judul, durasi, genre, tahun string) {
		film := map[string]string{
			"judul":  judul,
			"durasi": durasi,
			"genre":  genre,
			"tahun":  tahun,
		}
		*listFilm = append(*listFilm, film)
	}
}

func main() {
	//Soal 1
	var number = 30
	var numberPointer *int = &number

	var luasLingkaran float64 = 3.14 * float64(*numberPointer * *numberPointer)
	var kelilingLingkaran float64 = 2 * 3.14 * float64(*numberPointer)

	fmt.Printf("Luas Lingkaran: %f\n", luasLingkaran)
	fmt.Printf("Keliling Lingkaran: %f\n", kelilingLingkaran)
	fmt.Println("")

	//Soal 2
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)
	fmt.Println("")

	//Soal 3
	buah()
	fmt.Println("")

	//Soal 4
	tambahDataFilm(&dataFilm)("LOTR", "2 jam", "action", "1999")
	tambahDataFilm(&dataFilm)("avenger", "2 jam", "action", "2019")
	tambahDataFilm(&dataFilm)("spiderman", "2 jam", "action", "2004")
	tambahDataFilm(&dataFilm)("juon", "2 jam", "horror", "2004")

	fmt.Println("Data Film:")
	for i, film := range dataFilm {
		fmt.Printf("%d. tittle: %s\n Durasi: %s\n Genre: %s\n Tahun: %s\n", i+1, film["judul"], film["durasi"], film["genre"], film["tahun"])
	}
}
