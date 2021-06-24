package main

import "oop/employee"

func main() {
	// e := employee.Employee{
	// 	FirstName: "John",
	// 	LastName:  "Wick",
	// 	Id:        12034,
	// 	JobName:   "Software Engineer",
	// }
	// e.Person()

	e := employee.New("Faruqi", "Hafiz", 1222, "Cloud Engineer")
	e.Person()
	e2 := employee.New("Jhon", "Wick", 1022, "Software Engineer")
	e2.Person()
}
