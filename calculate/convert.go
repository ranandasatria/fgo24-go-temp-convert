package calculate

import (
	"fmt"
	"fgo24-go-temp-convert/history"
)

func CtoK(c int) {
	convert := c + 273
	result := fmt.Sprintf("%d°C = %d°K", c, convert)
	fmt.Println("Hasil konversi ke Kelvin:", convert, "°K")
	history.Add(result)
}

func CtoF(c int) {
	convert := (9.0 / 5.0) * float64(c) + 32
	result := fmt.Sprintf("%d°C = %.2f°F", c, convert)
	fmt.Printf("Hasil konversi ke Fahrenheit: %.2f°F\n", convert)
	history.Add(result)
}

func CtoR(c int) {
	convert := (4.0 / 5.0) * float64(c)
	result := fmt.Sprintf("%d°C = %.2f°R", c, convert)
	fmt.Printf("Hasil konversi ke Reamur: %.2f°R\n", convert)
	history.Add(result)
}
