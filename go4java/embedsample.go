package main

import "fmt"

type Person struct{ Name string }

func (p Person) Introduce() { fmt.Println("Hi, I'm", p.Name) }

type Employee struct {
	Person
	EmployeeID int
}

func main() {
	var e Employee
	e.Name = "Peter"
	e.EmployeeID = 1234

	e.Introduce()
}
