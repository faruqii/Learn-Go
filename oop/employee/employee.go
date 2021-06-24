package employee

import "fmt"

type Employee struct {
	FirstName string
	LastName string
	Id       int
	JobName  string
}

func New(FirstName string, LastName string, Id int, JobName string) Employee {
	e := Employee{FirstName, LastName, Id, JobName}
	return e
}

func (e Employee) Person() {
	fmt.Printf("%s %s ID is: %d\nThe Job Name is: %s\n", e.FirstName, e.LastName, e.Id, e.JobName)
}
