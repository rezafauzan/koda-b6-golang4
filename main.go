package main

import (
	"fmt"
	"os"
)

func main() {
	var temp float64
	for true {
		var pilihan int
		fmt.Println("Masukan suhu dalam bentuk celcius")
		fmt.Scanln(&temp)
		fmt.Println("1.Konversi ke fahrenheit")
		fmt.Println("2.Konversi ke kelvin")
		fmt.Println("3.Konversi ke reamur")
		fmt.Println("0.Keluar Dari Program")
		fmt.Scanln(&pilihan)
		if pilihan == 1 {
			fmt.Printf("Konversi suhu %.1f celcius ke fahrenheit adalah %.1f \n", temp, ((temp * 9 / 5) + 32))
		}
		if pilihan == 2 {
			fmt.Printf("Konversi suhu %.1f celcius ke fahrenheit adalah %.1f \n", temp, (temp + 273.15))
		}
		if pilihan == 3 {
			fmt.Printf("Konversi suhu %.1f celcius ke fahrenheit adalah %.1f \n", temp, (temp * 4 / 5))
		}
		if pilihan == 0 {
			fmt.Println("Pilihan tidak ada !")
			os.Exit(0)
		}
	}
}
