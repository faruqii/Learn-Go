package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	if b == 0 {
		panic("Cannot divide number with zero! ")
	}
	return a / b
}

func main() {
	var menu, a, b int
	fmt.Printf("Calculator\n1.Sum\n2.Subtract\n3.Multiply\n4.Divide\n")
	fmt.Scan(&menu)

	switch menu {

	case 1:
		fmt.Print("Enter First Number: ")
		fmt.Scan(&a)
		fmt.Print("Enter Second Number: ")
		fmt.Scan(&b)
		fmt.Print(sum(a, b))

	case 2:
		fmt.Print("Enter First Number: ")
		fmt.Scan(&a)
		fmt.Print("Enter Second Number: ")
		fmt.Scan(&b)
		fmt.Print(subtract(a, b))

	case 3:
		fmt.Print("Enter First Number: ")
		fmt.Scan(&a)
		fmt.Print("Enter Second Number: ")
		fmt.Scan(&b)
		fmt.Print(multiply(a, b))

	case 4:
		fmt.Print("Enter First Number: ")
		fmt.Scan(&a)
		fmt.Print("Enter Second Number: ")
		fmt.Scan(&b)
		fmt.Print(divide(a, b))

	default:
		fmt.Print("Wrong Menu!! ")
	}
}
