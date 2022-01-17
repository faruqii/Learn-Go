package main

import "fmt"

type Person struct {
	name   string
	age    int
	gender string
}

type Student struct {
	Person
	class string
	grade int
}

func main() {
	s1 := Student{}
	s1.name = "Faruqi"
	s1.age = 20
	s1.gender = "man"
	s1.class = "Second Year"
	s1.grade = 100

	fmt.Println(s1.name, s1.age, s1.gender, s1.class, s1.grade)
}