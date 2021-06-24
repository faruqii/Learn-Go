package main

import "fmt"

func main() {

	mymap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4}

	for key, val := range mymap {
		fmt.Printf("key %s has value %d\n", key, val)
	}
}
