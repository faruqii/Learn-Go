package main

import (
	"fmt"
)

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for idx, numbers := range numbers {
		fmt.Printf("Value at Index %d is %d\n", idx, numbers)
	}
}
