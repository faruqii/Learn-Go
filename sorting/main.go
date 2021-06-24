package main

import (
	"fmt"
)

func bublesort(items []int) []int {
	L := len(items)

	for i := 0; i < L; i++ {
		for j := 0; j < (L - 1 - i); j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
		
	} 
	return items
}

func main() {
	items := []int {10,2,3,13,4,23}
	items2 := []int {10,2,3,99,22,44}
	fmt.Println(bublesort(items))
	fmt.Print(bublesort(items2))
}