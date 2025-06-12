package main

import (
	"fgo24-go-temp-convert/calculate"
	"fgo24-go-temp-convert/history"
	"fmt"
	"os"
)

func main() {
	for{fmt.Println("Selamat datang\n1. Konversi Suhu\n2. Riwayat Konversi\n0. Keluar")
	var choice int
	fmt.Print("Masukkan pilihan: ")
	fmt.Scanln(&choice)
	var celcius int
	if choice == 1 {
		fmt.Print("Masukkan suhu dalam Celcius: ")
		fmt.Scanln(&celcius)
		calculate.CtoK(celcius)
		calculate.CtoF(celcius)
		calculate.CtoR(celcius)
	} else if choice == 2 {
		history.Show()
	} else if choice == 0 {
		os.Exit(0)
	} else {
		fmt.Println("Menu belum tersedia")
	}
	}
}
