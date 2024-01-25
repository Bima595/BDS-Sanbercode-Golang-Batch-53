package Models

import(
	"fmt"
)

//Soal 1
type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

//Soal 2
type Hp interface{
	Show()
}

//Soal 3
func LuasPersegi(sisi int, tampilkanKalimat bool) interface{} {
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