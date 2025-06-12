package calculate

import "fmt"

func CtoK(c int) {
	convert := c + 273
	fmt.Printf("Hasil konversi ke Kelvin: %d°K\n", convert)
}

func CtoF(c int) {
	convert := (9.0/5.0)*float64(c) + 32
	fmt.Printf("Hasil konversi ke Fahrenheit: %.2f°F\n", convert)
}

func CtoR(c int) {
	convert := (4.0 / 5.0) * float64(c)
	fmt.Printf("Hasil konversi ke Reamur: %.2f°R\n", convert)
}
