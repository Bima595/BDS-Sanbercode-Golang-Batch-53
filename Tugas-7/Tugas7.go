package main

import (
	"fmt"
)

//soal 1
type buah struct {
	nama string
	biji string
	warna string 
	harga int 
}

//soal 2
type segitiga struct {
	alas, tinggi int
}

func (s segitiga) Luas() float64 {
	return 0.5 * float64(s.alas) * float64(s.tinggi)
}

type persegi struct {
	sisi int
}

func (p persegi) Luas() int {
	return p.sisi * p.sisi
}

type persegiPanjang struct {
	panjang, lebar int
}

func (pp persegiPanjang) Luas() int {
	return pp.panjang * pp.lebar
}

//soal 3
type phone struct {
	name   string
	brand  string
	year   int
	colors *[]string
}

func (p *phone) tambahWarna(warna string) {
	*p.colors = append(*p.colors, warna)
}

func displayPhoneInfo(p *phone) {
	fmt.Printf("Name: %s\nBrand: %s\nYear: %d\nColors: %v\n\n", p.name, p.brand, p.year, *p.colors)
}

//soal 4
type movie struct {
	title    string
	duration int
	genre    string
	year     int
}

var dataFilm = []movie{}

func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie) {
	filmBaru := movie{
		title:    title,
		duration: duration,
		genre:    genre,
		year:     year,
	}
	*dataFilm = append(*dataFilm, filmBaru)
}

func tampilkanDataFilm(dataFilm []movie) {
	fmt.Println("Data Film:")
	for i, film := range dataFilm {
		fmt.Printf("%d. Title: %s\n   Durasi: %d jam\n   Genre: %s\n   Tahun: %d\n", i+1, film.title, film.duration, film.genre, film.year)
	}
} 

func main() {
//soal 1
	var nanas,jeruk,semangka,pisang buah 
	nanas.nama = "Nanas"
	nanas.biji = "Tidak"
	nanas.warna = "Kuning"
	nanas.harga = 9000

	jeruk.nama = "Jeruk"
	jeruk.biji = "Ada"
	jeruk.warna = "Oranye"
	jeruk.harga = 8000
	
	semangka.nama = "Semangka"
	semangka.biji = "Ada"
	semangka.warna = "Hijau & Merah"
	semangka.harga = 10000

	pisang.nama = "Pisang"
	pisang.biji = "Tidak"
	pisang.warna = "Kuning"
	pisang.harga = 5000

	fmt.Println("nama\t"+"\t"+"Warna\t"+"\t"+"Ada Bijinya\t"+"Harga")
	fmt.Println(nanas.nama + "\t" +"\t"+ nanas.warna+"\t"+"\t"+ nanas.biji +"\t"+"\t"+fmt.Sprintf("%d", nanas.harga))
	fmt.Println(jeruk.nama + "\t" +"\t"+ jeruk.warna+"\t"+"\t"+ jeruk.biji +"\t"+"\t"+fmt.Sprintf("%d", jeruk.harga))
	fmt.Println(semangka.nama + "\t"+ semangka.warna+"\t"+ semangka.biji +"\t"+"\t"+fmt.Sprintf("%d", semangka.harga))
	fmt.Println(pisang.nama + "\t" +"\t"+ pisang.warna+"\t"+"\t"+ pisang.biji +"\t"+"\t"+fmt.Sprintf("%d", pisang.harga))

//Soal 2
var Segitiga segitiga
	Segitiga.alas = 30
	Segitiga.tinggi = 50
	fmt.Println("")
	fmt.Printf("Luas Segitiga: %f\n", Segitiga.Luas())

	var persegirms persegi
	persegirms.sisi = 5
	fmt.Printf("Luas Persegi: %d\n", persegirms.Luas())

	var PersegiPanjang persegiPanjang
	PersegiPanjang.lebar = 30
	PersegiPanjang.panjang = 50
	fmt.Printf("Luas Persegi Panjang: %d\n", PersegiPanjang.Luas())

	//Soal 3
	oldBrandColors := []string{"Black", "White"}
	newBrandColors := []string{"Red", "Green", "Yellow"}

	var PhoneOldBrand phone
	PhoneOldBrand.name = "Iphone"
	PhoneOldBrand.brand = "XR"
	PhoneOldBrand.year = 2010
	PhoneOldBrand.colors = &oldBrandColors

	var PhoneNewBrand phone
	PhoneNewBrand.name = "Iphone"
	PhoneNewBrand.brand = "LatestModel"
	PhoneNewBrand.year = 2022
	PhoneNewBrand.colors = &newBrandColors

	fmt.Println("\nXr Old Brand:")
	displayPhoneInfo(&PhoneOldBrand)
	fmt.Println("Adding Blue to XR Old Brand Colors:")
	PhoneOldBrand.tambahWarna("Blue")
	displayPhoneInfo(&PhoneOldBrand)

	fmt.Println("\nXR New Brand:")
	displayPhoneInfo(&PhoneNewBrand)
	fmt.Println("Adding Purple to XR New Brand Colors:")
	PhoneNewBrand.tambahWarna("Purple")
	displayPhoneInfo(&PhoneNewBrand)

//Soal 4
	tambahDataFilm("LOTR", 120, "Action", 1999, &dataFilm)
	tambahDataFilm("Avenger", 120, "Action", 2019, &dataFilm)
	tambahDataFilm("Spiderman", 120, "Action", 2004, &dataFilm)
	tambahDataFilm("Juon", 120, "Horror", 2004, &dataFilm)

	tampilkanDataFilm(dataFilm)


}

