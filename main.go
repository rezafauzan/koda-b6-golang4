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
		if(pilihan == 1){
			fmt.Printf("Konversi suhu %d celcius ke fahrenheit adalah %d", temp, ((temp * 9/5)))
		}
	}
}