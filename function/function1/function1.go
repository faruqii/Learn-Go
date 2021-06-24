package main

import "fmt"

func celana(bahan string) {
	Loop:
		for {
			switch bahan {
			
			case "Abaka":
				fmt.Println("Keluarlah Celana Tahan Air")
				break Loop
			
			case "Pandan":
				fmt.Println("Keluarlah Celana Wangi")
				break Loop
		
			case "Batu":
				fmt.Println("Keluarlah Celana dengan perlindungan maksimal")
				break Loop
		
			default:
				fmt.Println("Bahan Tidak Bisa di Buat! ")
				break Loop
		
			}
		}
}

func main() {
	var bahan string
	fmt.Print("Masukkan Bahan: ")
	fmt.Scan(&bahan)
	celana(bahan)
}
