package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// soal 1
		var kalimat1 ="Bootcamp "
		var kalimat2 ="Digital "
		var kalimat3 ="Skill "
		var kalimat4 ="Sanbercode "
		var kalimat5 ="Golang"
		var FullKalimat = kalimat1 + kalimat2 + kalimat3 + kalimat4 + kalimat5
		fmt.Println(FullKalimat)

	// soal 2
		var halo = "Halo Dunia"
		var index = "Dunia"
		var ganti = "Golang"
		var kalimat = strings.Replace(halo, index, ganti, 1)
		fmt.Println("")	
		fmt.Println(kalimat)

	// soal 3
		var kataPertama = "saya";
		var kataKedua = "senang";
		var kataKetiga = "belajar";
		var find = "r"
		var gantiw = "R"
		var kataKeempat = "golang";
		var ubah3 = strings.Replace(kataKetiga,find, gantiw,1)
		var gede = strings.ToUpper(kataKeempat)
		fmt.Println("")
		fmt.Println(kataPertama, kataKedua, ubah3, gede)

	// soal 4
		var angkaPertama= "8";
		var angkaKedua= "5";
		var angkaKetiga= "6";
		var angkaKeempat = "7";
		intAngka1, _ := strconv.Atoi(angkaPertama)
		intAngka2, _ := strconv.Atoi(angkaKedua)
		intAngka3, _ := strconv.Atoi(angkaKetiga)
		intAngka4, _ := strconv.Atoi(angkaKeempat)
		var hitung = (intAngka1 + intAngka2 + intAngka3 + intAngka4)
		fmt.Println("")
		fmt.Println(hitung)	

	// soal 5
		kalimat = "halo halo bandung"
		find = "halo"
		ganti = "HI"
		var ubah = strings.Replace(kalimat, find, ganti, 2)

		angka := 2021
		ubah1 := strconv.Itoa(angka)
		fmt.Println("")
		fmt.Printf("%s - %s", ubah, ubah1)

}