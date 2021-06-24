package main

import (
	"fmt"
)

func main() {

	// arrays
	var ArrTwoD [4][5]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			ArrTwoD[i][j] = i * j
		}
	}
	fmt.Print(ArrTwoD)
}
