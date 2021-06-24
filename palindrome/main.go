package main

import "fmt"

func main() {
	var num, remainder, temp int
	var reverse int = 0

	fmt.Print("Enter the the number: ")
	fmt.Scan(&num)

	temp = num

	for {
		remainder = num % 10
		reverse = reverse*10 + remainder
		num /= 10

		if (num == 0) {
			break
		}
	}

	if (temp==reverse) {
		fmt.Printf("%d is a palindrome",temp)

	} else {
		fmt.Printf("%d is not a palindrome",temp)
	}
}