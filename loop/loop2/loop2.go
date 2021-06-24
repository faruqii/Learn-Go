package main

import "fmt"

func main() {

	var a int
	fmt.Print("Masukkan Angka")
	fmt.Scan(&a)

	for a > 0 {
		fmt.Println(a)
		a--
		if a == 1 {
			fmt.Print("its enough")
			break
		}
	}
}

