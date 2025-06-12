package main

import (
	"fgo24-go-temp-convert/calculate"
	"fmt"
)


func main(){
	fmt.Println("Selamat datang\n1. Konversi Suhu\n2. Riwayat Konversi")
	var choice int
	fmt.Print("Masukkan pilihan: ")
	fmt.Scanln(&choice)
	var celcius int
	if choice == 1{
		fmt.Print("Masukkan suhu dalam Celcius: ")
		fmt.Scanln(&celcius)
		calculate.CtoK(celcius)
		calculate.CtoF(celcius)
		calculate.CtoR(celcius)
	} else if choice == 2 {
		fmt.Println("Menu belum tersedia")
	} else {
		fmt.Println("Menu belum tersedia")
	}
}

