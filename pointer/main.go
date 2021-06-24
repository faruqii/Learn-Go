package main

import "fmt"

func change_value(val int) {
	val = 200

}

func change_reference(val *int) {
	*val = 500
}

func main() {

	var a int = 10

	fmt.Println(a)
	change_value(a)
	fmt.Println(a)
}
