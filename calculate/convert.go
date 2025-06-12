package calculate

import "fmt"

func CtoK(c int) {
	convert := c + 273
	fmt.Printf("Hasil konversi ke Kelvin adalah: %d\n",convert)
}

func CtoF(c int){
	convert := (9/5*c) + 32
	fmt.Printf("Hasil konversi ke Kelvin adalah: %d\n",convert)
}

func CtoR(c int){
	convert := (4/5) * c
	fmt.Printf("Hasil konversi ke Reamur adalah: %d\n",convert)

}