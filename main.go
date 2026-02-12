package main

import "fmt"

func main() {
	var temp int
	for true {
		var pilihan int
		fmt.Println("Masukan suhu dalam bentuk celcius")
		fmt.Scanln(&temp)
		fmt.Println("1.Konversi ke fahrenheit")
		fmt.Println("2.Konversi ke kelvin")
		fmt.Println("3.Konversi ke reamur")
		fmt.Println("0.Konversi ke reamur")
		fmt.Scanln(&pilihan)
	}
}