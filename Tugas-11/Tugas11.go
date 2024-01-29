package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
)

//Soal 1
var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
func displaySortedData(data []string, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Strings(data)
	for i, phone := range data {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(1 * time.Second)
	}
}

//Soal 2
var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}
func getMovies(ch chan string, movies ...string) {
	defer close(ch)
	for i, movie := range movies {
		ch <- fmt.Sprintf("%d. %s", i+1, movie)
	}
}

//Soal 3
type Result struct {
	JariJari    float64
	LuasLingkaran   float64
	KelilingLingkaran float64
	VolumeTabung   float64
}

func hitungLuasLingkaran(jariJari float64, ch chan Result, wag *sync.WaitGroup) {
	defer wag.Done()

	luas := math.Pi * jariJari * jariJari
	result := Result{
		JariJari:        jariJari,
		LuasLingkaran:   luas,
		KelilingLingkaran: 0,
		VolumeTabung:   0,
	}

	ch <- result
}

func hitungKelilingLingkaran(jariJari float64, ch chan Result, wag *sync.WaitGroup) {
	defer wag.Done()

	keliling := 2 * math.Pi * jariJari
	result := Result{
		JariJari:        jariJari,
		LuasLingkaran:   0,
		KelilingLingkaran: keliling,
		VolumeTabung:   0,
	}

	ch <- result
}

func hitungVolumeTabung(jariJari, tinggi float64, ch chan Result, wag *sync.WaitGroup) {
	defer wag.Done()

	volume := math.Pi * jariJari * jariJari * tinggi
	result := Result{
		JariJari:        jariJari,
		LuasLingkaran:   0,
		KelilingLingkaran: 0,
		VolumeTabung:   volume,
	}

	ch <- result
}

//Soal 4
type Hasil struct {
	Luas, Keliling, Volume float64
}

func hitungLuasPanjang(panjang, lebar float64, ch chan Hasil, wg2 *sync.WaitGroup) {
	defer wg2.Done()

	luas := panjang * lebar
	result := Hasil{
		Luas:     luas,
		Keliling: 0,
		Volume:   0,
	}

	ch <- result
}

func hitungKelilingPanjang(panjang, lebar float64, ch chan Hasil, wg2 *sync.WaitGroup) {
	defer wg2.Done()

	keliling := 2 * (panjang + lebar)
	result := Hasil{
		Luas:     0,
		Keliling: keliling,
		Volume:   0,
	}

	ch <- result
}

func hitungVolumeBalok(panjang, lebar, tinggi float64, ch chan Hasil, wg2 *sync.WaitGroup) {
	defer wg2.Done()

	volume := panjang * lebar * tinggi
	result := Hasil{
		Luas:     0,
		Keliling: 0,
		Volume:   volume,
	}

	ch <- result
}

func main() {
	//Soal 1
	var wg sync.WaitGroup
	wg.Add(1)
	go displaySortedData(phones, &wg)
	wg.Wait()

	//Soal 2
	moviesChannel := make(chan string)
	go getMovies(moviesChannel, movies...)
	fmt.Println("")
	fmt.Println("List Movies: ")
	for value := range moviesChannel {
  	fmt.Println(value)
	}

	//Soal 3
	var wag sync.WaitGroup
	ch := make(chan Result, 3)

	jariJari := []float64{8, 14, 20}
	tinggiTabung := 10

	for _, r := range jariJari {
		wag.Add(1)
		go hitungLuasLingkaran(r, ch, &wag)

		wag.Add(1)
		go hitungKelilingLingkaran(r, ch, &wag)

		wag.Add(1)
		go hitungVolumeTabung(r, float64(tinggiTabung), ch, &wag)
	}

	go func() {
		wag.Wait()
		close(ch)
	}()

	for hasil := range ch {
		fmt.Println("")
		fmt.Printf("Jari-Jari: %.2f\n", hasil.JariJari)
		fmt.Printf("Luas Lingkaran: %.2f\n", hasil.LuasLingkaran)
		fmt.Printf("Keliling Lingkaran: %.2f\n", hasil.KelilingLingkaran)
		fmt.Printf("Volume Tabung: %.2f\n", hasil.VolumeTabung)
		fmt.Println()
	}

	//Soal 4
	var wg2 sync.WaitGroup
	ca := make(chan Hasil, 3)

	panjang := 5.0
	lebar := 3.0
	tinggi := 2.0

	wg2.Add(1)
	go hitungLuasPanjang(panjang, lebar, ca, &wg2)

	wg2.Add(1)
	go hitungKelilingPanjang(panjang, lebar, ca, &wg2)

	wg2.Add(1)
	go hitungVolumeBalok(panjang, lebar, tinggi, ca, &wg2)

	go func() {
		wg2.Wait()
		close(ca)
	}()

	for {
		select {
		case result, ok := <-ca:
			if !ok {
				return
			}
			fmt.Println("")
			fmt.Printf("Luas Persegi Panjang: %.2f\n", result.Luas)
			fmt.Printf("Keliling Persegi Panjang: %.2f\n", result.Keliling)
			fmt.Printf("Volume Balok: %.2f\n", result.Volume)
			fmt.Println()
		}
	}
}
