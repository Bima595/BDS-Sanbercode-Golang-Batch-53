package Models

import(
	"fmt"
	"math"
)

//Soal 1
func (s SegitigaSamaSisi) Luas() int {
	return (s.Alas * s.Tinggi) / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return 3 * s.Alas
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

func (t Tabung) Volume() float64 {
	return math.Pi * math.Pow(t.JariJari, 2) * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * t.JariJari * (t.JariJari + t.Tinggi)
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return 2 * (float64(b.Panjang*b.Lebar) + float64(b.Panjang*b.Tinggi) + float64(b.Lebar*b.Tinggi))
}

//Soal 2
func (p Phone)Show(){
	fmt.Printf("Name: %s\nBrand: %s\nYear: %d\nColors: %v\n", p.Name, p.Brand, p.Year, p.Colors)
}

