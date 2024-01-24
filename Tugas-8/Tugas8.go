package main

import (
	"fmt"
	"math"
)

//Soal 1
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() int {
	return 3 * s.alas
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * (float64(b.panjang*b.lebar) + float64(b.panjang*b.tinggi) + float64(b.lebar*b.tinggi))
}

//Soal 2
type hp interface{
	Show()
}

type phone struct{
	name, brand string
	year int
	colors []string
 }

func (p phone)show(){
	fmt.Printf("Name: %s\nBrand: %s\nYear: %d\nColors: %v\n", p.name, p.brand, p.year, p.colors)
}

//Soal 3
func luasPersegi(sisi int, tampilkanKalimat bool) interface{} {
	if sisi == 0 {
		if tampilkanKalimat {
			return "Maaf anda belum menginput sisi dari persegi"
		}
		return nil
	}

	luas := sisi * sisi

	if tampilkanKalimat {
		return fmt.Sprintf("Luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	}

	return luas
}

func main() {
	//Soal 1
	segitiga := segitigaSamaSisi{alas: 4, tinggi: 5}
	fmt.Println("Segitiga Sama Sisi:")
	fmt.Println("Luas:", segitiga.luas())
	fmt.Println("Keliling:", segitiga.keliling())
	fmt.Println()

	persegi := persegiPanjang{panjang: 6, lebar: 8}
	fmt.Println("Persegi Panjang:")
	fmt.Println("Luas:", persegi.luas())
	fmt.Println("Keliling:", persegi.keliling())
	fmt.Println()

	silinder := tabung{jariJari: 3.0, tinggi: 7.0}
	fmt.Println("Tabung:")
	fmt.Println("Volume:", silinder.volume())
	fmt.Println("Luas Permukaan:", silinder.luasPermukaan())
	fmt.Println()

	kotak := balok{panjang: 3, lebar: 4, tinggi: 5}
	fmt.Println("Balok:")
	fmt.Println("Volume:", kotak.volume())
	fmt.Println("Luas Permukaan:", kotak.luasPermukaan())

	//Soal 2
	handphones := phone{
		name:"Samsung Galaxy Note 20",
		brand:"Samsung Galaxy Note 20",
		year:2020,
		colors:[]string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}
	
	fmt.Println(" ")
	handphones.show()
	
	//Soal 3
	fmt.Println(" ")
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	//Soal 4
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	angkaPertama := kumpulanAngkaPertama.([]int)
	angkaKedua := kumpulanAngkaKedua.([]int)

	total := 0
	for _, angka := range append(angkaPertama, angkaKedua...) {
		total += angka
	}

	prefixString := prefix.(string)

	fmt.Println("")
	output := fmt.Sprintf("%s%d + %d + %d + %d = %d", prefixString, angkaPertama[0], angkaPertama[1], angkaKedua[0], angkaKedua[1], total)
	fmt.Println(output)

}
