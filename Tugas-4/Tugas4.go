package main

import(
	"fmt"
)

func main() {
	//Soal 1
	for i := 1; i <= 20; i++ {
		if i%2 == 0{
			fmt.Println(i, "santai")
		} else{
			fmt.Println(i, "Berkualitas")
		}

		if i%3 == 0 && i%2 != 0 {
			fmt.Println(i, "I Love Coding")
		}
	}

	//Soal 2
	for i := 0; i <= 6; i++ {
		for j:= 1; j < i ; j++ {
			fmt.Printf("#")
		} 
		fmt.Println()
	}

	//Soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
		fmt.Println("")
		fmt.Println(kalimat[2:7])



	//Soal 4
	var sayuran = []string{"Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun"}

		for i, sayuran := range sayuran {
			fmt.Println("")
			fmt.Printf("%d. %s",i+1,sayuran)
		}
		fmt.Println("")

	//Soal 5
		var satuan = map[string]int{
			"panjang": 7,
			"lebar":   4,
			"tinggi":  6,
		  }
		var volbal = satuan["panjang"]*satuan["lebar"]*satuan["tinggi"]
		fmt.Println("")
		fmt.Println("Panjang = ",satuan["panjang"])
		fmt.Println("Lebar = ",satuan["lebar"])
		fmt.Println("Tinggi = ",satuan["tinggi"])  
		fmt.Println("Volume Balok = ",volbal)
}