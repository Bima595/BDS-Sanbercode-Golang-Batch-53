package main

import (
  "fmt"
  "errors"
  "sort"
  "time"
  "math"
  "flag"
)

//Soal 1
func printDetails(kalimat string, tahun int) {
	fmt.Printf("Kalimat: %s\n", kalimat)
	fmt.Printf("Tahun: %d\n", tahun)
}

//Soal 2
func kelilingSegitigaSamaSisi(sisi int, tampilkanDetail bool) (string, error) {
	if sisi == 0 {
		return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	}

	keliling := sisi * 3

	if tampilkanDetail {
		return fmt.Sprintf("Keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, keliling), nil
	}

	return fmt.Sprintf("%d", keliling), nil
}

//Soal 3
func tambahAngka(nilai int, total *int) {
	*total += nilai
}

func cetakAngka(total *int) {
	fmt.Printf("Total angka: %d\n", *total)
}

//Soal 4
var phones = []string{}

func tambahPhone(brand string, data *[]string) {
	*data = append(*data, brand)
}

//Soal 5
func luasLingkaran(jariJari float64) int {
	luas := math.Pi * math.Pow(jariJari, 2)
	return int(math.Round(luas))
}

func kelilingLingkaran(jariJari float64) int {
	keliling := 2 * math.Pi * jariJari
	return int(math.Round(keliling))
}

//soal 6
func luasPersegiPanjang(panjang, lebar float64) float64 {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang, lebar float64) float64 {
	return 2 * (panjang + lebar)
}


func main() {
	//Soal 1
	kalimat := "Golang Backend Development"
	tahun := 2021
	defer printDetails(kalimat, tahun)
	fmt.Println("Ini akan dieksekusi sebelum defer")

	//Soal 2
	result1, err1 := kelilingSegitigaSamaSisi(4, true)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println(result1)
	}

	result2, err2 := kelilingSegitigaSamaSisi(8, false)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println(result2)
	}

	result3, err3 := kelilingSegitigaSamaSisi(0, true)
	if err3 != nil {
		fmt.Println("Error:", err3)
	} else {
		fmt.Println(result3)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	_, err4 := kelilingSegitigaSamaSisi(0, false)
	if err4 != nil {
		fmt.Println("Error:", err4)
	}

	//Soal 3 
	angka := 1
	defer cetakAngka(&angka)

	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	//Soal 4
	tambahPhone("Xiaomi", &phones)
	tambahPhone("Asus", &phones)
	tambahPhone("IPhone", &phones)
	tambahPhone("Samsung", &phones)
	tambahPhone("Oppo", &phones)
	tambahPhone("Realme", &phones)
	tambahPhone("Vivo", &phones)

	sort.Strings(phones)

	for i, phone := range phones {
		time.Sleep(1 * time.Second)
		fmt.Printf("%d. %s\n", i+1, phone)
	}

	//Soal 5
	jariJari1 := 7.0
	jariJari2 := 10.0
	jariJari3 := 15.0

	luas1 := luasLingkaran(jariJari1)
	keliling1 := kelilingLingkaran(jariJari1)

	luas2 := luasLingkaran(jariJari2)
	keliling2 := kelilingLingkaran(jariJari2)

	luas3 := luasLingkaran(jariJari3)
	keliling3 := kelilingLingkaran(jariJari3)

	fmt.Printf("Jari-jari: %.1f\n", jariJari1)
	fmt.Printf("Luas Lingkaran: %d\n", luas1)
	fmt.Printf("Keliling Lingkaran: %d\n", keliling1)

	fmt.Printf("\nJari-jari: %.1f\n", jariJari2)
	fmt.Printf("Luas Lingkaran: %d\n", luas2)
	fmt.Printf("Keliling Lingkaran: %d\n", keliling2)

	fmt.Printf("\nJari-jari: %.1f\n", jariJari3)
	fmt.Printf("Luas Lingkaran: %d\n", luas3)
	fmt.Printf("Keliling Lingkaran: %d\n", keliling3)

	//Soal 6
	var panjang, lebar float64
	flag.Float64Var(&panjang, "panjang", 0, "Panjang persegi panjang")
	flag.Float64Var(&lebar, "lebar", 0, "Lebar persegi panjang")
	flag.Parse()

	if panjang <= 0 || lebar <= 0 {
		fmt.Println("Panjang dan lebar harus lebih dari 0.")
		return
	}

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)

	fmt.Printf("Panjang: %.2f\n", panjang)
	fmt.Printf("Lebar: %.2f\n", lebar)
	fmt.Printf("Luas Persegi Panjang: %.2f\n", luas)
	fmt.Printf("Keliling Persegi Panjang: %.2f\n", keliling)
}