package history

import "fmt"

var Records []string 

func Add(record string) {
	Records = append(Records, record)
}

func Show() {
	if len(Records) == 0 {
		fmt.Println("Belum ada riwayat konversi.")
		return
	}
	fmt.Println("Riwayat Konversi:")
	for i, r := range Records {
		fmt.Printf("%d. %s\n", i+1, r)
	}
}
