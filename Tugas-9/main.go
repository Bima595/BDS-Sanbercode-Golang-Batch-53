package main

import (
	"BDS-Sanbercode-Golang-Batch-53/Tugas-9/Models"
	"fmt"
)

func main() {
	//Soal 1
	segitiga := Models.SegitigaSamaSisi{Alas: 4, Tinggi: 5}
	fmt.Println("Segitiga Sama Sisi:")
	fmt.Println("Luas:", segitiga.Luas())
	fmt.Println("Keliling:", segitiga.Keliling())
	fmt.Println()

	persegi := Models.PersegiPanjang{Panjang: 6, Lebar: 8}
	fmt.Println("Persegi Panjang:")
	fmt.Println("Luas:", persegi.Luas())
	fmt.Println("Keliling:", persegi.Keliling())
	fmt.Println()

	silinder := Models.Tabung{JariJari: 3.0, Tinggi: 7.0}
	fmt.Println("Tabung:")
	fmt.Println("Volume:", silinder.Volume())
	fmt.Println("Luas Permukaan:", silinder.LuasPermukaan())
	fmt.Println()

	kotak := Models.Balok{Panjang: 3, Lebar: 4, Tinggi: 5}
	fmt.Println("Balok:")
	fmt.Println("Volume:", kotak.Volume())
	fmt.Println("Luas Permukaan:", kotak.LuasPermukaan())

	//Soal 2
	handphones := Models.Phone{
		Name:"Samsung Galaxy Note 20",
		Brand:"Samsung Galaxy Note 20",
		Year:2020,
		Colors:[]string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}
	
	fmt.Println(" ")
	handphones.Show()
	
	//Soal 3
	fmt.Println(" ")
	fmt.Println(Models.LuasPersegi(4, true))
	fmt.Println(Models.LuasPersegi(8, false))
	fmt.Println(Models.LuasPersegi(0, true))
	fmt.Println(Models.LuasPersegi(0, false))

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
