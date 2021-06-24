package main

import "fmt"

func main() {

	var number int

	fmt.Print("Enter Number: ")
	_, err := fmt.Scan(&number)

	if err != nil {
		fmt.Printf("Can Assign String Type! ")

	} else {
		fmt.Println(number)
	}

}
