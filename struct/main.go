package main

import "fmt"

type Person struct {
	name, class, address string
	age, id              int
}

func main() {
	
	p1 := Person{"Cassandra","SI-44-04","Jakarta",19,1204}
	fmt.Println(p1)
}
